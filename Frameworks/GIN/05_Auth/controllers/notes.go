package controllers

import (
	"ginLearning/05_Auth/middleware"
	"ginLearning/05_Auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	notesService *services.NotesService
}

func (n *NotesController) Init(notesService *services.NotesService) {
	n.notesService = notesService
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine) {
	notes := router.Group("/notes")
	notes.Use(middleware.CheckMiddleware)
	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNotes())
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	note, err := n.notesService.GetNotesService()
	if err != nil {
		return func(c *gin.Context) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"note": note,
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {
	type NoteBody struct {
		Title  string `json:"title"`
		Status bool   `json:"status"`
	}

	return func(c *gin.Context) {
		var noteBody NoteBody
		if err := c.BindJSON(&noteBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		note, err := n.notesService.CreateNotesService(noteBody.Title, noteBody.Status)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"note": note,
		})
	}
}
