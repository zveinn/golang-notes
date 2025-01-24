package main

import (
	"log/slog"
	"os"
	"time"
)

func structuredLogging() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	logger.Debug("This is a debug message", slog.String("key", "value"))
	logger.Info("This is an info message", slog.Int("count", 42))
	logger.Warn("This is a warning", slog.Float64("temperature", 98.6))
	logger.Error("This is an error", "err", "disk full")

	logger.Info("User login",
		slog.String("user", "john_doe"),
		slog.String("action", "login"),
		slog.Time("timestamp", time.Now()),
	)

	logger.Info("Request processed",
		slog.Group("request",
			slog.String("method", "GET"),
			slog.String("path", "/users"),
		),
		slog.Int("status", 200),
		slog.String("equal_sign_as_parameter", "==x==1=="),
	)
}
