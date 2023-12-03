package usecase

import (
	"github.com/Shabashkin93/warning_tracker/internal/domain/status"
)

type Status interface {
	GetAll() (dto *status.Status, code int, err error)
}
