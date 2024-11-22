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
