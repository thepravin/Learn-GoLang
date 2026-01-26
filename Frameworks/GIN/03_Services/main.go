package main

import (
	"ginLearning/03_Services/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	notesController := &controllers.NotesController{}
	notesController.InitNotesControllerRoutes(router)

	router.Run(":8080")
}
