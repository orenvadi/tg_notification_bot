package notif_bot

import (
	"log"
	"strings"

	"github.com/orenvadi/tg_notification_bot/internal/storage"
	"github.com/tucnak/telebot"
)

// State constants for the state machine
const (
	StateIdle = iota
	StateWaitingForEmail
	StateWaitingForPassword
)

// UserState stores the current state of the user in the bot conversation
type UserState struct {
	State  int
	Email  string
	UserID int64
}

var userStates = make(map[int64]*UserState) // Map Telegram user ID to their state

func StartBot(bot *telebot.Bot, db *storage.DB) {
	bot.Handle(telebot.OnText, func(m *telebot.Message) {
		userID := m.Sender.ID

		// Get user state or initialize it
		state, exists := userStates[int64(userID)]
		if !exists {
			state = &UserState{State: StateIdle, UserID: int64(userID)}
			userStates[int64(userID)] = state
		}

		switch state.State {
		case StateIdle:
			handleIdleState(bot, state, m)
		case StateWaitingForEmail:
			handleEmailInput(bot, state, m)
		case StateWaitingForPassword:
			handlePasswordInput(bot, state, m, db)
		default:
			_, err := bot.Send(m.Sender, "Something went wrong. Please type /start to reset.")
			if err != nil {
				log.Printf("bot could not send message, err: %v", err)
			}
			state.State = StateIdle
		}
	})

	bot.Start()
}

func handleIdleState(bot *telebot.Bot, state *UserState, m *telebot.Message) {
	if strings.ToLower(m.Text) == "/start" {
		_, err := bot.Send(m.Sender, "Welcome! Please enter your email address to log in.")
		if err != nil {
			log.Printf("bot could not send message, err: %v", err)
		}
		state.State = StateWaitingForEmail
	} else {
		_, err := bot.Send(m.Sender, "Type /start to begin.")
		if err != nil {
			log.Printf("bot could not send message, err: %v", err)
		}
	}
}

func handleEmailInput(bot *telebot.Bot, state *UserState, m *telebot.Message) {
	state.Email = strings.TrimSpace(m.Text)
	if !strings.Contains(state.Email, "@") {
		_, err := bot.Send(m.Sender, "Invalid email address. Please enter a valid email.")
		if err != nil {
			log.Printf("bot could not send message, err: %v", err)
		}
		return
	}
	_, err := bot.Send(m.Sender, "Great! Now, please enter your password.")
	if err != nil {
		log.Printf("bot could not send message, err: %v", err)
	}
	state.State = StateWaitingForPassword
}

func handlePasswordInput(bot *telebot.Bot, state *UserState, m *telebot.Message, db *storage.DB) {
	password := strings.TrimSpace(m.Text)

	// Mock authentication with MDelivery backend
	if authenticateUser(state.Email, password) {
		err := db.SaveUser(state.Email, state.UserID)
		if err != nil {
			_, err := bot.Send(m.Sender, "Failed to save your account. Please try again.")
			if err != nil {
				log.Printf("bot could not send message, err: %v", err)
			}
			log.Printf("Database error: %v", err)
		} else {
			_, err := bot.Send(m.Sender, "Login successful! You will now receive notifications here.")
			if err != nil {
				log.Printf("bot could not send message, err: %v", err)
			}
			state.State = StateIdle
		}
	} else {
		_, err := bot.Send(m.Sender, "Invalid email or password. Please type /start to try again.")
		if err != nil {
			log.Printf("bot could not send message, err: %v", err)
		}
		state.State = StateIdle
	}
}

func authenticateUser(email, password string) bool {
	// Replace with actual MDelivery authentication API call
	// Mocking the authentication process
	return email == "test@example.com" && password == "password"
}
