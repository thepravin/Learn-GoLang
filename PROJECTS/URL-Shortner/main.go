package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"urlShortner/models"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var baseURL string

func initDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	baseURL = os.Getenv("BASE_URL")

	// Ensure your DSN is correct and the database service is running
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database...")
	}

	db = database
	err = db.AutoMigrate(&models.URL{})
	if err != nil {
		log.Fatal("Failed to auto-migrate schema : ", err)
	}

	log.Println("Database connection established successfully.....")
}

func GenerateRandomString(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)[:n]
}

func shortenURL(c echo.Context) error {
	type Request struct {
		URL string `json:"url"`
	}

	req := new(Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid Request"})
	}

	shortCode := GenerateRandomString(6)

	url := models.URL{
		OriginalURL: req.URL,
		ShortCode:   shortCode,
	}

	if err := db.Create(&url).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not save the URL"})
	}

	return c.JSON(http.StatusOK, echo.Map{"short_url": fmt.Sprintf("%s/%s", baseURL, shortCode)})
}

func redirectURL(c echo.Context) error {
	code := c.Param("code")
	var url models.URL

	if err := db.Where("short_code=?", code).First(&url).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Url not found"})
	}
	db.Model(&url).Update("clicks", url.Clicks+1)

	return c.Redirect(http.StatusFound, url.OriginalURL)
}

func stats(c echo.Context) error {
	code := c.Param("code")
	var url models.URL

	if err := db.Where("short_code=?", code).First(&url).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Url not found"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"original_url": url.OriginalURL,
		"short_code":   url.ShortCode,
		"clicks":       url.Clicks,
	})
}

func main() {
	initDB()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/shorten", shortenURL)
	e.GET("/:code", redirectURL)
	e.GET("/stats/:code", stats)

	e.Logger.Fatal(e.Start(":8090"))
}
