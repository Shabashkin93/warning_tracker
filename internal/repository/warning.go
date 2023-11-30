package repository

import "github.com/Shabashkin93/warning_tracker/internal/domain/warning"

type Warning interface {
	Create(in *warning.WarningCreate) (id string, err error)
}
