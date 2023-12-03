package usecase

import (
	"github.com/Shabashkin93/warning_tracker/internal/domain/warning"
)

type Warning interface {
	Create(in *warning.WarningCreate) (result warning.WarningResponse, err error)
}
