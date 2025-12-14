package main

import (
	"chatApp/internal/api"
	"chatApp/internal/model"
	"chatApp/internal/store"
	"chatApp/internal/ws"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	store.ConnectDatabase()
	store.DB.AutoMigrate(&ws.OutboundMessage{})
	store.DB.AutoMigrate(&model.Message{})
	hub := ws.NewHub()

	go hub.Run()

	e := echo.New()

	api.RegisterRoutes(e, hub)

	log.Println("Server is running on : 8090")
	e.Logger.Fatal(e.Start(":8090"))
}
