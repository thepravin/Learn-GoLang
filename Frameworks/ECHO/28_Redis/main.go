package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	rdb *redis.Client
	ctx = context.Background()
)

type User struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

func initDB() *gorm.DB {
	// dsn := "user:password@tcp(host:port)/database"
	dsn := "app_user:app_pass@tcp(127.0.0.1:3606)/demo?parseTime=true"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect the database....")
	}

	err = database.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to auto-migrate schema : ", err)
	}

	log.Println("Database connection established and schema migrated successfully.")

	return database
}

func initRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect with redis %v", err)
	}
	return rdb
}

// func getAllUsers(c echo.Context) error {
// 	val, err := rdb.Get(ctx, "all_users").Result()
// 	if err == redis.Nil {
// 		var users []User
// 		if err := db.Find(&users).Error; err != nil {
// 			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "DB Error"})
// 		}
// 		data, _ := json.Marshal(users)
// 		rdb.Set(ctx, "all_users", data, 10*time.Minute)

// 		return c.JSON(http.StatusOK, users)
// 	} else if err != nil {
// 		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Redis Error"})
// 	}

// 	var users []User
// 	if err := json.Unmarshal([]byte(val), &users); err != nil {
// 		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to parse the data"})
// 	}

// 	return c.JSON(http.StatusOK, users)
// }

func getAllUsers(c echo.Context) error {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "DB Error"})
	}

	data, _ := json.Marshal(users)
	rdb.Set(ctx, "all_users", data, 10*time.Minute)

	return c.JSON(http.StatusOK, users)
}

func createUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	var lastUser User
	if err := db.Order("id desc").First(&lastUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "DB Error"})
	}

	u.ID = lastUser.ID + 1
	if err := db.Create(&u).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "DB Error"})
	}

	var users []User
	if err := db.Find(&users).Error; err != nil {
		data, _ := json.Marshal(users)
		rdb.Set(ctx, "all_users", data, 10*time.Minute)
	}
	data, _ := json.Marshal(u)
	rdb.Set(ctx, fmt.Sprintf("user:%d", u.ID), data, 10*time.Minute)
	return c.JSON(http.StatusCreated, u)
}

func main() {

	db = initDB()
	rdb = initRedis()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/allusers", getAllUsers)
	e.POST("/createuser", createUser)

	e.Logger.Fatal(e.Start(":8090"))

}
