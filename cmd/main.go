package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"message_service/config"
	db "message_service/internal/database"
	api "message_service/internal/handlers"
	"message_service/internal/kafka"
)

func main() {
	app := fiber.New()

	cfg := config.LoadConfig()

	database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer database.Close()

	producer, err := kafka.NewProducer(cfg)
	if err != nil {
		log.Fatalf("failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	api.SetupRoutes(app, database, producer)

	log.Fatal(app.Listen(":8080"))
}
