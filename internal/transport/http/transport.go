package transport

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/Shabashkin93/warning_tracker/internal/config"
	"github.com/Shabashkin93/warning_tracker/internal/logging"
	"github.com/Shabashkin93/warning_tracker/internal/service"
	server "github.com/Shabashkin93/warning_tracker/internal/transport/http/http_gin"
	"github.com/Shabashkin93/warning_tracker/internal/transport/http/http_gin/status"
	"github.com/Shabashkin93/warning_tracker/internal/transport/http/http_gin/warning"
)

const versionApi = "/v1"

type Transport struct {
	logger *logging.LoggerEntry
	Status
	Warning
	handler http.Handler
	server  *http.Server
	ctx     context.Context
}

func NewTransport(ctx context.Context, service *service.Service, logger *logging.LoggerEntry, cfg *config.Config) *Transport {
	handler := server.Init(logger)
	server := &http.Server{
		Handler: handler,
	}
	return &Transport{
		logger:  logger,
		Status:  status.NewTransport(service, handler, cfg.HTTP.URL.Status, cfg.IsDebug),
		Warning: warning.NewTransport(service, handler, cfg.HTTP.URL.Warning, cfg.IsDebug),
		handler: handler,
		server:  server,
		ctx:     ctx,
	}
}

func (t *Transport) StartServer(cfg *config.Config) {
	addr := ":" + cfg.HTTP.Port
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		t.logger.Fatal(t.ctx, "Error start web server", fmt.Sprintf("%v", err))

	}

	t.Status.Register(versionApi, t.handler)
	t.Warning.Register(versionApi, t.handler)

	server := &http.Server{
		Handler: t.handler,
	}

	go func() {
		server.Serve(listener)
	}()
	t.logger.Info(t.ctx, "Started server", slog.String("address", addr))

}

func (t *Transport) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := t.server.Shutdown(ctx); err != nil {
		t.logger.Error(t.ctx, "Could not shut down server correctly", slog.String("error", fmt.Sprintf("%v", err)))

	}
}
