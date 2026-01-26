package main

import (
	"ginLearning/05_Auth/controllers"
	"ginLearning/05_Auth/db"
	"ginLearning/05_Auth/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := db.InitDB()

	if db == nil {
		log.Fatal("DB not connected.....")
	}

	authService := services.InitAuthService(db)
	authController := controllers.InitAuthController(authService)
	authController.InitAuthControllerRoutes(router)

	notesService := &services.NotesService{}
	notesService.InitService(db)

	notesController := &controllers.NotesController{}
	notesController.Init(notesService)
	notesController.InitNotesControllerRoutes(router)

	router.Run(":7757")
}
