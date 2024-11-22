package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	DB *sql.DB
}

func InitDB() (*DB, error) {
	db, err := sql.Open("sqlite3", "./bot.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		email TEXT PRIMARY KEY,
		telegram_user_id INTEGER
	)`)
	if err != nil {
		return nil, err
	}

	return &DB{
		DB: db,
	}, nil
}

func (db *DB) Close() error {
	return db.DB.Close()
}
