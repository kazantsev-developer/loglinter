// Package positive provides valid logging test cases for false positive verification
package positive

import "log/slog"

// CheckPositive executes only valid logger calls that must pass without errors
func CheckPositive() {
	slog.Info("starting server on port 8080")
	slog.Error("failed to connect to database")
	slog.Warn("something went wrong")
	slog.Debug("api request completed")

	slog.Info("token validated")
	slog.Info("user authenticated successfully")
	slog.Info("server started")
	slog.Error("connection failed")
}
