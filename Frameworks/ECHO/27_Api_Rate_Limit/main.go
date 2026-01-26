package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserApiRate struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `json:"name"`
	APIKey    string    `gorm:"uniqueIndex;size:255" json:"api_key"`
	CreatedAT time.Time `json:"created_at"`
}

type RateLimit struct {
	Key         string `gorm:"primaryKey;size:255"` // Column is named 'key'
	Count       int
	WindowStart time.Time `gorm:"index"`
	CreatedAT   time.Time // Must be initialized to avoid '0000-00-00' error
	UpdateAt    time.Time // Must be initialized/updated
}

var db *gorm.DB

func initDB() {
	// Ensure your DSN is correct and the database service is running
	dsn := "app_user:app_pass@tcp(127.0.0.1:3606)/demo?parseTime=true"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database...")
	}

	db = database
	err = db.AutoMigrate(&UserApiRate{}, &RateLimit{})
	if err != nil {
		log.Fatal("Failed to auto-migrate schema : ", err)
	}

	log.Println("Database connection established successfully.....")
}

func generateAPIKey() string {
	bytes := make([]byte, 16)
	_, _ = rand.Read(bytes)
	return "API-" + hex.EncodeToString(bytes)
}

// Rate limiter main logic
func RateLimiter(limit int, window time.Duration, exceedStatus int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.Request().Header.Get("X-API-Key")
			if key == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"error": "Missing X-API-Key header",
				})
			}

			// 1. Validate the API key exists in the UserApiRate table
			var user UserApiRate
			if err := db.First(&user, "api_key = ?", key).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return c.JSON(http.StatusUnauthorized, echo.Map{
						"error": "Invalid API key",
					})
				}
				log.Println("RateLimiter: User API Key lookup failed:", err)
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"error": "Database lookup failed during authentication",
				})
			}

			// 2. Rate Limiting logic
			now := time.Now().UTC()
			var rl RateLimit

			// Try to find the existing rate limit record
			err := db.First(&rl, "`key` = ?", key).Error

			if err != nil && err != gorm.ErrRecordNotFound {
				log.Println("RateLimiter: Database error during RateLimit lookup:", err)
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"error": "Database error during rate limiting",
				})
			}

			// If record not found, or window has expired, reset the count
			if err == gorm.ErrRecordNotFound || now.Sub(rl.WindowStart) >= window {
				rl = RateLimit{
					Key:         key,
					Count:       1,
					WindowStart: now,
					CreatedAT:   now, // FIX: Initialize to current time
					UpdateAt:    now, // FIX: Initialize to current time
				}
				db.Save(&rl)
				return next(c)
			}

			// Limit exceeded
			if rl.Count >= limit {
				retryAfter := int(window.Seconds()) - int(now.Sub(rl.WindowStart).Seconds())
				c.Response().Header().Set("Retry-After", fmt.Sprint(retryAfter))
				return c.JSON(exceedStatus, echo.Map{
					"error":       "rate limit exceeded",
					"retry_after": retryAfter,
					"status":      exceedStatus,
				})
			}

			// Increment count and save
			rl.Count++
			rl.UpdateAt = now // FIX: Update the UpdateAt timestamp
			db.Save(&rl)
			return next(c)
		}
	}
}

func main() {
	initDB()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/signup", func(c echo.Context) error {
		type SignupRequest struct {
			Name string `json:"name"`
		}
		req := new(SignupRequest)

		if err := c.Bind(req); err != nil || strings.TrimSpace(req.Name) == "" {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Name is required or invalid"})
		}

		apiKey := generateAPIKey()
		user := UserApiRate{Name: req.Name, APIKey: apiKey, CreatedAT: time.Now()}

		result := db.Create(&user)
		if result.Error != nil {
			log.Println("Database Create Error:", result.Error)
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error":   "Failed to create user in database",
				"details": result.Error.Error(),
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "User created successfully",
			"api_key": apiKey,
		})
	})

	rateLimit := RateLimiter(15, 15*time.Second, http.StatusTooManyRequests)

	// Protected route uses the rate limiting middleware
	e.GET("/data", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Welcome to the protected data api",
			"time":    time.Now().Format(time.RFC3339),
		})
	}, rateLimit) // Apply middleware here

	e.Logger.Fatal(e.Start(":8090"))

}
