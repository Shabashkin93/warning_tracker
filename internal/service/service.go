package service

import (
	"context"

	"github.com/Shabashkin93/warning_tracker/internal/logging"
	"github.com/Shabashkin93/warning_tracker/internal/repository"
	"github.com/Shabashkin93/warning_tracker/internal/service/status"
	"github.com/Shabashkin93/warning_tracker/internal/service/warning"
	"github.com/microcosm-cc/bluemonday"
)

type Service struct {
	Status
	Warning
}

func NewService(ctx context.Context, repos *repository.Repository, logger logging.Logger) *Service {
	sanity := bluemonday.UGCPolicy()
	return &Service{
		Status:  status.NewService(repos),
		Warning: warning.NewService(ctx, sanity, repos, logger),
	}
}
