package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/segmentio/kafka-go"
	"message_service/internal/models"
	"message_service/internal/services"
)

func SetupRoutes(app *fiber.App, db *pgxpool.Pool, producer *kafka.Writer) {
	app.Post("/message", func(c *fiber.Ctx) error {
		var message models.Message
		if err := c.BodyParser(&message); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
		}

		if err := services.SaveMessage(db, message); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot save message"})
		}

		if err := producer.WriteMessages(c.Context(), kafka.Message{
			Value: []byte(message.Content),
		}); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot send message to Kafka"})
		}

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": "message received"})
	})

	app.Get("/statistics", func(c *fiber.Ctx) error {
		count, err := services.GetProcessedMessagesCount(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot get statistics"})
		}
		return c.JSON(fiber.Map{"processed_messages_count": count})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Message Service is running!")
	})
}
