package logging

import (
	"context"
	"os"

	logger "github.com/Shabashkin93/warning_tracker/internal/logging/slog"
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
	Handler interface{}
	LogFile *os.File
}

func (logger *LoggerEntry) Stop() {
	logger.LogFile.Close()
}

func NewLogger() *LoggerEntry {
	logger, handler, logfile := logger.NewLogger()
	return &LoggerEntry{
		Logger:  logger,
		Handler: handler,
		LogFile: logfile,
	}
}
