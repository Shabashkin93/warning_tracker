package empty_log

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

var l *slog.Logger

var logfile *os.File

type logger struct {
	logger *slog.Logger
}

func NewLogger(logLevel string) (*logger, *os.File) {
	loggerEntry := getLogger()
	return &logger{
		logger: loggerEntry,
	}, logfile
}

func getLogger() *slog.Logger {
	return l
}

func (l *logger) Panic(ctx context.Context, format string, v ...interface{}) {
	panic(fmt.Sprintf(format, v...))
}

func (l *logger) Fatal(ctx context.Context, format string, v ...interface{}) {
	os.Exit(1)
}

func (l *logger) Error(ctx context.Context, format string, v ...interface{}) {
}

func (l *logger) Warn(ctx context.Context, format string, v ...interface{}) {
}

func (l *logger) Info(ctx context.Context, format string, v ...interface{}) {
}

func (l *logger) Debug(ctx context.Context, format string, v ...interface{}) {
}
