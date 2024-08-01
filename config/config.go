package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DatabaseURLA string
	KafkaURL     string
	KafkaTopic   string
}

func LoadConfig() *Config {
	err := godotenv.Load("./.env", ".env")
	if err != nil {
		log.Println("не удалось спарсить переменное окружение", err)
	}
	log.Println(os.Getenv("DATABASE_URL"))
	return &Config{
		DatabaseURLA: os.Getenv("DATABASE_URLA"),
		KafkaURL:     os.Getenv("KAFKA_URL"),
		KafkaTopic:   os.Getenv("KAFKA_TOPIC"),
	}
}
