package main

import (
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/db"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/http/handler"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/repository"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/router"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/ws"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot find env file")
	}

	r := router.New()

	dbConnection, err := db.New()
	if err != nil {
		log.Fatal("Cannot connect to database.")
	}

	db.AutoMigrate(dbConnection)

	v1 := r.Group("/api")

	ur := repository.NewUserRepo(dbConnection)
	cr := repository.NewContactRepo(dbConnection)
	chr := repository.NewChatRepo(dbConnection)
	mr := repository.NewMessageRepo(dbConnection)

	h := handler.NewHandler(ur, cr, chr, mr)
	h.Register(v1)

	// Handle WebSocket connections
	hub := ws.NewHub()
	go hub.Run()
	r.GET("/ws", func(c echo.Context) error {
		ws.ServeWs(hub, c.Response(), c.Request())
		return nil
	})

	if err := r.Start("0.0.0.0:8080"); err != nil {
		log.Fatalf("server failed to start %v", err)
	}
}
