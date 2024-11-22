# Telegram Notification Bot API

## Send a notification to a user

### POST `/api/send_notification`

Send a notification to a user by their email.

#### Request Body
```json
{
  "email": "string",
  "message": "string"
}
```

#### Responses

| Status Code | Description                       | Response Body        |
|-------------|-----------------------------------|----------------------|
| 200         | Notification sent successfully   | `"Notification sent successfully"` |
| 400         | Invalid request body             | `"Invalid request body"`          |
| 404         | User not found                   | `"User not found"`                |
| 500         | Internal server error            | `"Internal server error"`         |





## How to build and run

1. Clone this repository:
   ```bash
   git clone https://github.com/your-username/tg_notification_bot.git
   cd tg_notification_bot
   ```

2. Create a `.env` file in the project root with the following content:
   ```env
   TELEGRAM_BOT_TOKEN=your-telegram-bot-token
   ```

   Replace `your-telegram-bot-token` with your actual Telegram bot token from [BotFather](https://core.telegram.org/bots#botfather).

---

### Build and Run

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Build the project:
   ```bash
   go build -o telegram_bot cmd/bot/main.go
   ```

3. Run the application:
   ```bash
   ./telegram_bot
   ```
