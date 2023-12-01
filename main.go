package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shabashkin93/warning_tracker/internal/config"
	"github.com/Shabashkin93/warning_tracker/internal/logging"
	"github.com/Shabashkin93/warning_tracker/internal/repository"
	db "github.com/Shabashkin93/warning_tracker/internal/repository/postgres"
	"github.com/Shabashkin93/warning_tracker/internal/service"
	transport "github.com/Shabashkin93/warning_tracker/internal/transport/http"
)

func main() {

	var ctx = context.Background()

	logger := logging.NewLogger()
	defer logger.Stop()

	cfg := config.GetConfig(ctx, logger)

	database, err := db.Initialize(ctx, logger, cfg)
	if err != nil {
		logger.Fatal(ctx, "Could not set up database", slog.String("error", fmt.Sprintf("%v", err)))
	}

	repos := repository.NewRepository(ctx, logger, cfg, &database)
	defer repos.Stop()

	services := service.NewService(ctx, repos, logger)

	transport := transport.NewTransport(ctx, services, logger, cfg)
	transport.StartServer(cfg)
	defer transport.Stop()

	// listen for ctrl+c signal from terminal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	logger.Info(ctx, fmt.Sprint(<-ch))
	logger.Info(ctx, "Stopping API server.")
	ctx.Done()
}
