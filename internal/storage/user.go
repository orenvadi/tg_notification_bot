package storage

import (
	"database/sql"
	"fmt"

	"github.com/orenvadi/tg_notification_bot/domain/models"
)

func (db *DB) SaveUser(email string, telegramID int64) error {
	// Use INSERT OR REPLACE to either insert a new user or update an existing one
	query := `
		INSERT OR REPLACE INTO users (email, telegram_user_id)
		VALUES (?, ?)
	`
	_, err := db.DB.Exec(query, email, telegramID)
	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

func (db *DB) GetUserByEmail(email string) (*models.User, error) {
	const op = "db.sqlite.getUserByEmail"
	row := db.DB.QueryRow("SELECT email, telegram_user_id FROM users WHERE email = ?", email)

	var user models.User
	err := row.Scan(&user.Email, &user.TelegramUserID)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("op=%s, User not found err: %v", op, err) // User not found
	}
	if err != nil {
		return nil, fmt.Errorf("op=%s, DB error err: %v", op, err) // User not found
	}

	return &user, nil
}
