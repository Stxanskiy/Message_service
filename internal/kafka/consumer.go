package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"message_service/config"
)

func NewConsumer(cfg *config.Config) (*kafka.Reader, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.KafkaURL},
		Topic:   cfg.KafkaTopic,
		GroupID: "message-consumer-group",
	})
	return reader, nil
}

func StartConsumer(cfg *config.Config, processMessage func([]byte) error) {
	reader, err := NewConsumer(cfg)
	if err != nil {
		log.Fatalf("failed to create Kafka consumer: %v", err)
	}
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error while reading message: %v", err)
			continue
		}

		if err := processMessage(msg.Value); err != nil {
			log.Printf("error while processing message: %v", err)
		}
	}
}
