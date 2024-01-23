package logging

import (
	"context"
)

type Logger interface {
	Panic(ctx context.Context, format string, v ...interface{})
	Fatal(ctx context.Context, format string, v ...interface{})
	Error(ctx context.Context, format string, v ...interface{})
	Warn(ctx context.Context, format string, v ...interface{})
	Info(ctx context.Context, format string, v ...interface{})
	Debug(ctx context.Context, format string, v ...interface{})
}

type LoggerEntry struct {
	Logger
}

// New logger instance
func NewLogger(logger Logger) *LoggerEntry {
	return &LoggerEntry{
		Logger: logger,
	}
}
