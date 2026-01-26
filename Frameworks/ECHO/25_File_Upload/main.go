package main

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type FileUpload struct {
	gorm.Model
	FileName string `json:"file_name,omitempty"`
	FileType string `json:"file_type,omitempty"`
	FileData []byte `json:"-" gorm:"type:longblob"`
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

	err = db.AutoMigrate(&FileUpload{})
	if err != nil {
		log.Fatal("Failed to auto-migrate schema : ", err)
	}

	log.Println("Database connection established and schema migrated successfully.")
}

func uploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "File is required"})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot open the file"})
	}
	defer src.Close()

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cnannot read the file"})
	}

	uploadFile := FileUpload{
		FileName: file.Filename,
		FileType: file.Header.Get("Content-Type"),
		FileData: fileBytes,
	}
	if err := db.Create(&uploadFile).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to save the file"})
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message":   "File saved successfully",
			"id":        uploadFile.ID,
			"file_name": uploadFile.FileName,
			"file_type": uploadFile.FileType})

}

func getFile(c echo.Context) error {
	id := c.Param("id")
	var file FileUpload

	if err := db.First(&file, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "File not found"})
	}

	return c.Stream(http.StatusOK, file.FileType, bytes.NewReader(file.FileData))
}

func main() {
	initDB()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/upload", uploadFile)
	e.GET("/file/:id", getFile)

	e.Logger.Fatal(e.Start(":8090"))
}
