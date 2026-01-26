package main

import (
	"fmt"
	"ginLearning/04_DB/controllers"
	"ginLearning/04_DB/database"
	"ginLearning/04_DB/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := database.InitDB()

	if db == nil {
		fmt.Println("Error while connecting DB..")
	}

	notesService := &services.NotesService{}
	notesService.InitService(db)

	notesController := &controllers.NotesController{}
	notesController.Init(notesService)
	notesController.InitNotesControllerRoutes(router)

	router.Run(":7757")
}
