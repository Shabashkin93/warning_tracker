package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shabashkin93/warning_tracker/internal/config"
	"github.com/Shabashkin93/warning_tracker/internal/logging"
	"github.com/Shabashkin93/warning_tracker/internal/repository"
	"github.com/Shabashkin93/warning_tracker/internal/service"
	transport "github.com/Shabashkin93/warning_tracker/internal/transport/http"
)

func main() {

	var ctx = context.Background()

	logger := logging.NewLogger()
	defer logger.Stop()

	cfg := config.GetConfig(ctx, logger)

	repos := repository.NewRepository(ctx, logger, cfg)
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
