package logger

import (
	"log/slog"
	"os"
)

const (
	varLocal string = "local"
	varDev string = "dev"
	varProd string = "prod"
)


type Logger struct {
	*slog.Logger
}


func NewLogger(env string, level string) *Logger {
	logLevel := new(slog.LevelVar)

	switch level {
	case "debug":
		logLevel.Set(slog.LevelDebug)
	case "info":
		logLevel.Set(slog.LevelInfo)
	case "warn":
		logLevel.Set(slog.LevelWarn)
	case "error":
		logLevel.Set(slog.LevelError)
	}

	var logger *slog.Logger
	switch env {
	case varLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case varDev:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	case varProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	}
	return &Logger{logger}
}