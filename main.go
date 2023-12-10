package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shabashkin93/warning_tracker/internal/config"
	"github.com/Shabashkin93/warning_tracker/internal/logging"
	logger "github.com/Shabashkin93/warning_tracker/internal/logging/slog"
	"github.com/Shabashkin93/warning_tracker/internal/repository"
	"github.com/Shabashkin93/warning_tracker/internal/repository/cache/redis"
	db "github.com/Shabashkin93/warning_tracker/internal/repository/postgres"
	transport "github.com/Shabashkin93/warning_tracker/internal/transport/http"
	"github.com/Shabashkin93/warning_tracker/internal/usecase"
)

func main() {

	var ctx = context.Background()

	cfg, err := config.GetConfig()
	if err != nil {
		slog.Error("Failed get config")
		os.Exit(1)
	}

	loggerInt, logfile := logger.NewLogger(cfg.LogLevel)

	logger := logging.NewLogger(loggerInt, logfile)
	defer logger.Stop()

	database, err := db.Initialize(ctx, logger, cfg)
	if err != nil {
		logger.Fatal(ctx, "Could not set up database", slog.String("error", fmt.Sprintf("%v", err)))
	}

	cache, err := redis.Init(ctx, logger, cfg.REDIS.Address, cfg.REDIS.Port, cfg.REDIS.Password, time.Duration(cfg.REDIS.Timeout))
	if err != nil {
		logger.Fatal(ctx, "Could not set up cache", slog.String("error", fmt.Sprintf("%v", err)))
	}

	repos := repository.NewRepository(ctx, logger, cfg, &database, &cache)
	defer repos.Stop()

	usecases := usecase.NewService(ctx, repos, logger)

	transport := transport.NewTransport(ctx, usecases, logger, cfg)
	transport.StartServer(cfg)
	defer transport.Stop()

	// listen for ctrl+c signal from terminal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	logger.Info(ctx, fmt.Sprint(<-ch))
	logger.Info(ctx, "Stopping API server.")
	ctx.Done()
}
