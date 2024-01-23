package status

import (
	"github.com/Shabashkin93/warning_tracker/internal/repository/postgres"
	"github.com/Shabashkin93/warning_tracker/pkg/logging"
)

type repository struct {
	logger   logging.Logger
	database *postgres.Database
}

func NewRepository(i interface{}, logger logging.Logger) *repository {
	database := i.(*postgres.Database)
	return &repository{
		database: database,
		logger:   logger,
	}
}

func (r *repository) GetStatus() (bool, error) {

	err := r.database.Conn.Ping()
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}
