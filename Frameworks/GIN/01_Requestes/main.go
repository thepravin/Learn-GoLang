package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/num/:id/:newId", func(ctx *gin.Context) {
		var id = ctx.Param("id")
		var newId = ctx.Param("newId")

		ctx.JSON(http.StatusOK, gin.H{
			"user_id":     id,
			"user_new_id": newId,
		})
	})

	/*
	   router.POST("/login", func(ctx *gin.Context) {
	   		type MeRequest struct {
	   			Email    string `json:"email"`
	   			Password string `json:"password"`
	   		}

	   		var meRequest MeRequest

	   		ctx.Bind(&meRequest)

	   		ctx.JSON(http.StatusOK, gin.H{
	   			"email":    meRequest.Email,
	   			"password": meRequest.Password,
	   		})
	   	})
	*/

	router.POST("/login", func(ctx *gin.Context) {
		type MeRequest struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password"`
		}

		var meRequest MeRequest

		if err := ctx.BindJSON(&meRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"email":    meRequest.Email,
			"password": meRequest.Password,
		})
	})

	//   router.Run() // listens on 0.0.0.0:8080 by default
	router.Run(":7757")
}
