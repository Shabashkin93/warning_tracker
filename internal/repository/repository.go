package repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Shabashkin93/warning_tracker/internal/config"
	"github.com/Shabashkin93/warning_tracker/internal/logging"
	"github.com/Shabashkin93/warning_tracker/internal/repository/cache"
	"github.com/Shabashkin93/warning_tracker/internal/repository/postgres/status"
	"github.com/Shabashkin93/warning_tracker/internal/repository/postgres/warning"
)

type Repository struct {
	logger   logging.Logger
	DataBase DataBase
	ctx      context.Context
	Cache
	Status
	Warning
}

func (r *Repository) Stop() {
	r.DataBase.Shutdown(r.ctx, r.logger)
	r.Cache.Shutdown()
}

func NewRepository(ctx context.Context, logger logging.Logger, cfg *config.Config, database DataBase) *Repository {
	var cacheEntry *cache.Cache
	cache, err := cache.Init(ctx, cfg.REDIS.Address, cfg.REDIS.Port, cfg.REDIS.Password, time.Duration(cfg.REDIS.Timeout))
	if err != nil {
		logger.Fatal(ctx, "Could not set up cache", slog.String("error", fmt.Sprintf("%v", err)))
	}
	cacheEntry = &cache

	return &Repository{
		logger:   logger,
		DataBase: database,
		ctx:      ctx,
		Cache:    cacheEntry,
		Status:   status.NewRepository(database, logger),
		Warning:  warning.NewRepository(ctx, database, cfg.DB.Schema+"."+cfg.DB.Table.Warning, logger, time.Duration(cfg.DB.Timeout)),
	}
}
