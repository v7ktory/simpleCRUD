package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/v7ktory/simpleCRUD/internal/repository"
	"github.com/v7ktory/simpleCRUD/internal/service"
	"github.com/v7ktory/simpleCRUD/internal/transport/rest"
	"github.com/v7ktory/simpleCRUD/pkg/database/postgres"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSL_MODE"),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := rest.NewHandler(services)

	handlers.InitRoutes().Run(os.Getenv("APP_PORT"))
}
