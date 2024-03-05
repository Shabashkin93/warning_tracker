package status

import (
	"github.com/Shabashkin93/warning_tracker/internal/repository/postgres"
	"github.com/Shabashkin93/warning_tracker/pkg/logging"
)

type repository struct {
	logger   logging.Logger
	database *postgres.Database
}

func NewRepository(i interface{}, logger logging.Logger) (repos *repository) {
	if i != nil {
		repos = &repository{}
		database := i.(*postgres.Database)
		repos.database = database
		repos.logger = logger
	}
	return
}

func (r *repository) GetStatus() (bool, error) {

	err := r.database.Conn.Ping()
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}
