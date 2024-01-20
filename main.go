package main

import (
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/db"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/http/handler"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/repository"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/router"
	"github.com/joho/godotenv"
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

	h := handler.NewHandler(ur)
	h.Register(v1)

	if err := r.Start("0.0.0.0:8080"); err != nil {
		log.Fatalf("server failed to start %v", err)
	}
}
