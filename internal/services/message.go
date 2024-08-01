package services

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"message_service/internal/models"
)

func SaveMessage(db *pgxpool.Pool, message models.Message) error {
	_, err := db.Exec(context.Background(), "INSERT INTO messages (content, processed) VALUES ($1, $2)", message.Content, false)
	return err
}

func MarkMessageProcessed(db *pgxpool.Pool, id int) error {
	_, err := db.Exec(context.Background(), "UPDATE messages SET processed = true WHERE id = $1", id)
	return err
}

func GetProcessedMessagesCount(db *pgxpool.Pool) (int, error) {
	var count int
	err := db.QueryRow(context.Background(), "SELECT COUNT(*) FROM messages WHERE processed = true").Scan(&count)
	return count, err
}
