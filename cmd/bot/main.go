package main

import (
	"log"
	"net/http"
	"os"

	notif_bot "github.com/orenvadi/tg_notification_bot/internal/bot"
	notifications_handler "github.com/orenvadi/tg_notification_bot/internal/handlers/notifications"

	"github.com/joho/godotenv"
	"github.com/orenvadi/tg_notification_bot/internal/storage"
	"github.com/tucnak/telebot"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the Telegram bot token from the environment
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatalf("TELEGRAM_BOT_TOKEN is not set in the environment")
	}

	// Initialize the database
	db, err := storage.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Create bot instance with poller
	bot, err := telebot.NewBot(telebot.Settings{
		Token: botToken,
		Poller: &telebot.LongPoller{
			Timeout: 10, // Adjust as needed (in seconds)
		},
	})
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Register bot handlers
	go notif_bot.StartBot(bot, db)

	// Register REST API routes
	http.HandleFunc("/api/send_notification", notifications_handler.New(bot, db))

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
