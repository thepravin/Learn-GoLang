package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	NAME        string  `json:"name"`
	CATEGORY    string  `json:"category"`
	PRICE       float64 `json:"price"`
	DESCRIPTION string  `json:"description"`
}

var db *gorm.DB

func initDB() {
	// dsn := "user:password@tcp(host:port)/database"
	dsn := "app_user:app_pass@tcp(127.0.0.1:3606)/demo?parseTime=true"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect the database....")
	}

	db = database

	err = db.AutoMigrate(&Product{})
	if err != nil {
		log.Fatal("Failed to auto-migrate schema : ", err)
	}

	log.Println("Database connection established and schema migrated successfully.")
}

func seedData(c echo.Context) error {
	sample := []Product{
		{
			NAME:        "Wireless Mechanical Keyboard",
			CATEGORY:    "Electronics",
			PRICE:       129.99,
			DESCRIPTION: "Compact 60% layout with tactile brown switches and RGB lighting.",
		},
		{
			NAME:        "Stainless Steel Water Bottle",
			CATEGORY:    "Home Goods",
			PRICE:       24.50,
			DESCRIPTION: "Insulated 32oz bottle keeps drinks cold for 24 hours.",
		},
		{
			NAME:        "Organic Coffee Beans",
			CATEGORY:    "Groceries",
			PRICE:       15.99,
			DESCRIPTION: "Medium roast, single-origin beans from Ethiopia (12 oz bag).",
		},
		{
			NAME:        "Noise-Cancelling Headphones",
			CATEGORY:    "Electronics",
			PRICE:       249.00,
			DESCRIPTION: "Over-ear design with industry-leading active noise cancellation.",
		},
		{
			NAME:        "Yoga Mat",
			CATEGORY:    "Fitness",
			PRICE:       35.75,
			DESCRIPTION: "Extra-thick, non-slip mat made from eco-friendly TPE material.",
		},
		{
			NAME:        "Fountain Pen Set",
			CATEGORY:    "Office Supplies",
			PRICE:       45.00,
			DESCRIPTION: "Elegant pen with a medium nib and five ink cartridges.",
		},
		{
			NAME:        "Portable Bluetooth Speaker",
			CATEGORY:    "Electronics",
			PRICE:       79.95,
			DESCRIPTION: "Waterproof speaker with 10 hours of playtime and deep bass.",
		},
		{
			NAME:        "Gardening Tool Set",
			CATEGORY:    "Home Goods",
			PRICE:       55.49,
			DESCRIPTION: "Includes trowel, transplanter, and cultivator with ergonomic handles.",
		},
		{
			NAME:        "Digital Drawing Tablet",
			CATEGORY:    "Electronics",
			PRICE:       99.00,
			DESCRIPTION: "10x6 inch active area tablet for digital art and graphic design.",
		},
		{
			NAME:        "Silk Pillowcase",
			CATEGORY:    "Beauty",
			PRICE:       39.99,
			DESCRIPTION: "100% pure mulberry silk for reduced hair frizz and skin creasing.",
		},
	}

	db.Create(&sample)
	return c.JSON(http.StatusOK, echo.Map{"message": "Product saved successfully"})
}

func getProducts(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")
	sortField := c.QueryParam("sortField")
	sortOrder := c.QueryParam("sortOrder")
	filter := c.QueryParam("filter")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 5
	}

	offset := (page - 1) * limit

	query := db.Model(&Product{})
	if filter != "" {
		filterPattern := "%" + strings.ToLower(filter) + "%"
		query = query.Where("LOWER(name) LIKE ? OR LOWER(category) LIKE ?", filterPattern, filterPattern)
	}

	if sortField != "" {
		order := "asc"
		if strings.ToLower(sortOrder) == "desc" {
			order = "desc"
		}
		query = query.Order(fmt.Sprintf("%s %s", sortField, order))
	}

	var total int64
	query.Count(&total)

	var product []Product
	if err := query.Limit(limit).Offset(offset).Find(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch data"})
	}

	totalPages := (int(total) + limit - 1) / limit

	return c.JSON(http.StatusOK, echo.Map{
		"page":        page,
		"limit":       limit,
		"total_items": total,
		"total_pages": totalPages,
		"data":        product,
	})

}

func main() {
	initDB()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/seeddata", seedData)
	e.GET("/prodcts", getProducts)

	e.Logger.Fatal(e.Start(":8090"))
}
