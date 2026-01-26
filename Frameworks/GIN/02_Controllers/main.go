package main

import (
	"ginLearning/02_Controllers/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	notesController := &controllers.NotesController{}
	notesController.InitNotesControllerRoutes(router)

	router.Run(":7757")
}
