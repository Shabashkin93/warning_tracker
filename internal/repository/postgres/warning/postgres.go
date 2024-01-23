package warning

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Shabashkin93/warning_tracker/internal/domain/warning"
	"github.com/Shabashkin93/warning_tracker/internal/repository/postgres"
	"github.com/Shabashkin93/warning_tracker/pkg/logging"
	"github.com/pkg/errors"
)

type repository struct {
	logger   logging.Logger
	database *postgres.Database
	table    *postgres.Table
	ctx      context.Context
	timeout  time.Duration
}

func NewRepository(ctx context.Context, i interface{}, tableName string, logger logging.Logger, timeout time.Duration) *repository {
	var err error
	database := i.(*postgres.Database)
	query := &postgres.Query{}

	query.Create, err = database.Conn.PrepareNamed(postgres.SanityQuery(fmt.Sprintf(`
	INSERT INTO %s (
			branch,
			commit,
			count,
			created_by,
			created_at
		) VALUES (
			:branch,
			:commit,
			:count,
			:created_by,
			:created_at
		)
		RETURNING
			id
		;
	`, tableName)))
	if err != nil {
		logger.Fatal(ctx, "prepared statement for \"Create\"", slog.String("err", err.Error()))
	}

	table := &postgres.Table{
		Name:  tableName,
		Query: query,
	}

	return &repository{
		database: database,
		logger:   logger,
		table:    table,
		ctx:      ctx,
		timeout:  timeout,
	}
}

func (r *repository) Create(in *warning.WarningCreate) (id string, err error) {
	if in == nil {
		err = errors.Errorf("repository/Create: Null input data")
		return
	}

	item := pgWarning{}
	domainToPgWarning(in, &item)

	ctx, cancel := context.WithTimeout(r.ctx, r.timeout*time.Second)
	defer cancel()

	err = r.table.Query.Create.GetContext(ctx, &item, &item)
	if err != nil {
		r.logger.Error(r.ctx, "Postgres req", slog.String("table", r.table.Name), slog.String("item", fmt.Sprintf("%v", item)), slog.String("error", fmt.Sprintf("%v", err)))

		err = errors.Errorf("repository/Create: Failed create in db Warning item")
		return
	}

	if item.ID.Valid {
		id = item.ID.String
	}

	return
}
