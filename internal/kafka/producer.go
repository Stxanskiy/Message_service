package kafka

import (
	"github.com/segmentio/kafka-go"
	"message_service/config"
)

func NewProducer(cfg *config.Config) (*kafka.Writer, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{cfg.KafkaURL},
		Topic:   cfg.KafkaTopic,
	})
	return writer, nil
}
