package slog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

var l *slog.Logger

var logfile *os.File

const (
	LevelDebug slog.Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

var levelNames = map[slog.Leveler]string{
	LevelDebug: "DEBUG",
	LevelInfo:  "INFO",
	LevelWarn:  "WARN",
	LevelError: "ERROR",
	LevelFatal: "FATAL",
	LevelPanic: "PANIC",
}

var levelValues = map[string]slog.Leveler{
	"DEBUG": LevelDebug,
	"INFO":  LevelInfo,
	"WARN":  LevelWarn,
	"ERROR": LevelError,
	"FATAL": LevelFatal,
	"PANIC": LevelPanic,
}

type logger struct {
	logger *slog.Logger
}

func NewLogger(logLevel string) (*logger, *os.File) {
	initLogger(logLevel)
	loggerEntry := getLogger()
	return &logger{
		logger: loggerEntry,
	}, logfile
}

func getLogger() *slog.Logger {
	return l
}

func initLogger(logLevel string) {
	levelVal, exists := levelValues[logLevel]
	if !exists {
		levelVal = LevelError
	}

	opts := &slog.HandlerOptions{
		Level: levelVal,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := levelNames[level]
				if !exists {
					levelLabel = level.String()
				}

				a.Value = slog.StringValue(levelLabel)
			}

			return a
		},
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	l = slog.New(handler)
	logfile = os.Stdout
}

func (l *logger) Panic(ctx context.Context, format string, v ...interface{}) {
	l.logger.Log(ctx, LevelPanic, format, v...)
	panic(fmt.Sprintf(format, v...))
}

func (l *logger) Fatal(ctx context.Context, format string, v ...interface{}) {
	l.logger.Log(ctx, LevelFatal, format, v...)
	os.Exit(1)
}

func (l *logger) Error(ctx context.Context, format string, v ...interface{}) {
	l.logger.Log(ctx, LevelError, format, v...)
}

func (l *logger) Warn(ctx context.Context, format string, v ...interface{}) {
	l.logger.Log(ctx, LevelWarn, format, v...)
}

func (l *logger) Info(ctx context.Context, format string, v ...interface{}) {
	l.logger.Log(ctx, LevelInfo, format, v...)
}

func (l *logger) Debug(ctx context.Context, format string, v ...interface{}) {
	l.logger.Log(ctx, LevelDebug, format, v...)
}
