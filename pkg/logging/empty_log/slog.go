package empty_log

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type logger struct {
	logger *slog.Logger
}

// New logger instance
func NewLogger(logLevel string, out *os.File) *logger {
	return &logger{
		logger: &slog.Logger{},
	}
}

// Mock for PANIC level log
func (l *logger) Panic(ctx context.Context, format string, v ...interface{}) {
	panic(fmt.Sprintf(format, v...))
}

// Mock for FATAL level log
func (l *logger) Fatal(ctx context.Context, format string, v ...interface{}) {
	os.Exit(1)
}

// Mock for ERROR level log
func (l *logger) Error(ctx context.Context, format string, v ...interface{}) {
}

// Mock for WARN level log
func (l *logger) Warn(ctx context.Context, format string, v ...interface{}) {
}

// Mock for INFO level log
func (l *logger) Info(ctx context.Context, format string, v ...interface{}) {
}

// Mock for DEBUG level log
func (l *logger) Debug(ctx context.Context, format string, v ...interface{}) {
}
