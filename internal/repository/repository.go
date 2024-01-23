package repository

import (
	"context"
	"time"

	"github.com/Shabashkin93/warning_tracker/internal/config"
	"github.com/Shabashkin93/warning_tracker/internal/repository/postgres/status"
	"github.com/Shabashkin93/warning_tracker/internal/repository/postgres/warning"
	"github.com/Shabashkin93/warning_tracker/pkg/logging"
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

func NewRepository(ctx context.Context, logger logging.Logger, cfg *config.Config, database DataBase, cache Cache) *Repository {
	return &Repository{
		logger:   logger,
		DataBase: database,
		ctx:      ctx,
		Cache:    cache,
		Status:   status.NewRepository(database, logger),
		Warning:  warning.NewRepository(ctx, database, cfg.DB.Schema+"."+cfg.DB.Table.Warning, logger, time.Duration(cfg.DB.Timeout)),
	}
}
