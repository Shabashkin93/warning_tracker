package logging

import (
	"context"
	"os"
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
	LogFile *os.File
}

func (logger *LoggerEntry) Stop() {
	logger.LogFile.Close()
}

func NewLogger(logger Logger, logfile *os.File) *LoggerEntry {
	return &LoggerEntry{
		Logger:  logger,
		LogFile: logfile,
	}
}
