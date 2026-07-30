// Package negative provides invalid logging test cases for verification
package negative

import "log/slog"

// CheckNegative executes invalid logger calls that must trigger linter errors
func CheckNegative() {
	slog.Info("Starting server on port 8080")   // want "log message must start with a lowercase letter"
	slog.Error("Failed to connect to database") // want "log message must start with a lowercase letter"

	slog.Info("запуск сервера")                    // want "log message must be in English only"
	slog.Error("ошибка подключения к базе данных") // want "log message must be in English only"

	slog.Info("server started!🚀")      // want "log message must not contain special characters or emojis"
	slog.Error("connection failed!!!") // want "log message must not contain special characters or emojis"

	slog.Info("user password: secret_val") // want "log message contains potentially sensitive data"
	slog.Debug("api_key=secret_key")       // want "log message contains potentially sensitive data"
}
