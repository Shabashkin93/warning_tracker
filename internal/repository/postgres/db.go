package postgres

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/Shabashkin93/warning_tracker/internal/config"
	"github.com/Shabashkin93/warning_tracker/pkg/logging"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var ErrNoMatch = errors.New("not found")

type Query struct {
	Create *sqlx.NamedStmt
	GetOne *sqlx.NamedStmt
}

type Table struct {
	Name  string
	Query *Query
}

type Database struct {
	Conn        *sqlx.DB
	LoggerEntry logging.Logger
	Table       *Table
}

func SanityQuery(query string) (cleanQuery string) {
	cleanQuery = strings.Replace(query, "	", "", -1)
	cleanQuery = strings.Replace(cleanQuery, "\n", " ", -1)

	return
}

func Initialize(ctx context.Context, logger logging.Logger, cfg *config.Config) (db Database, err error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Address, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Db, cfg.DB.SslMode)

	ctx, cancel := context.WithTimeout(ctx, time.Duration(cfg.DB.Timeout)*time.Second)
	defer cancel()

	conn, err := sqlx.ConnectContext(ctx, "pgx", dsn)
	if err != nil {
		logger.Fatal(ctx, "Unable to connect to database", slog.String("error", fmt.Sprintf("%v", err)))
	}

	db.LoggerEntry = logger
	db.Conn = conn

	db.LoggerEntry.Info(ctx, "Database connection established")
	return
}

func (db *Database) Shutdown(ctx context.Context, logger logging.Logger) {
	err := db.Conn.Close()
	if err != nil {
		logger.Error(ctx, "Failed disconnect to postgresql db")
	} else {
		logger.Info(ctx, "Success disconnect to postgresql db")
	}
}
