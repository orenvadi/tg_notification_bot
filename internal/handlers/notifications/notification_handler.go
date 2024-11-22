package notifications_handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/orenvadi/tg_notification_bot/domain/models"
	"github.com/tucnak/telebot"
)

type NotificationRequest struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

type UserGetter interface {
	GetUserByEmail(email string) (*models.User, error)
}

func New(bot *telebot.Bot, userGetter UserGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the method is POST
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Parse the request body
		var req NotificationRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate the request fields
		if req.Email == "" || req.Message == "" {
			http.Error(w, "Email and Message are required", http.StatusBadRequest)
			return
		}

		// Retrieve the user from the database
		user, err := userGetter.GetUserByEmail(req.Email)
		if err != nil {
			log.Printf("Database error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Check if the user exists
		if user == nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		// Send the notification via Telegram bot
		recipient := &telebot.User{ID: int(user.TelegramUserID)}
		if _, err := bot.Send(recipient, req.Message); err != nil {
			log.Printf("Failed to send message: %v", err)
			http.Error(w, "Failed to send notification", http.StatusInternalServerError)
			return
		}

		// Respond with success
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("Notification sent successfully"))
		if err != nil {
			log.Printf("could not write to http writer, err: %v", err)
		}
	}
}
