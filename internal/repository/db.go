package repository

import (
	"context"

	"github.com/Shabashkin93/warning_tracker/pkg/logging"
)

type DataBase interface {
	Shutdown(ctx context.Context, logger logging.Logger)
}
