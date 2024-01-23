package slog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

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

// New logger instance
func NewLogger(logLevel string, out *os.File) *logger {
	return &logger{
		logger: initLogger(logLevel, out),
	}
}

// Init slog ang get slog handler
func initLogger(logLevel string, out *os.File) (loggerEntry *slog.Logger) {
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

	handler := slog.NewJSONHandler(out, opts)
	loggerEntry = slog.New(handler)
	return
}

// Write PANIC level log and run panic
func (loggerEntry *logger) Panic(ctx context.Context, format string, v ...interface{}) {
	loggerEntry.logger.Log(ctx, LevelPanic, format, v...)
	panic(fmt.Sprintf(format, v...))
}

// Write FATAL level log and run exit
func (loggerEntry *logger) Fatal(ctx context.Context, format string, v ...interface{}) {
	loggerEntry.logger.Log(ctx, LevelFatal, format, v...)
	os.Exit(1)
}

// Write ERROR level log
func (loggerEntry *logger) Error(ctx context.Context, format string, v ...interface{}) {
	loggerEntry.logger.Log(ctx, LevelError, format, v...)
}

// Write WARN level log
func (loggerEntry *logger) Warn(ctx context.Context, format string, v ...interface{}) {
	loggerEntry.logger.Log(ctx, LevelWarn, format, v...)
}

// Write INFO level log
func (loggerEntry *logger) Info(ctx context.Context, format string, v ...interface{}) {
	loggerEntry.logger.Log(ctx, LevelInfo, format, v...)
}

// Write DEBUG level log
func (loggerEntry *logger) Debug(ctx context.Context, format string, v ...interface{}) {
	loggerEntry.logger.Log(ctx, LevelDebug, format, v...)
}
