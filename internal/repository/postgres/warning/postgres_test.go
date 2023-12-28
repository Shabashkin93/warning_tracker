package warning_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

var conn *sqlx.DB
var IPAddress string

const (
	POSTGRES_USER     = "wtrack"
	POSTGRES_PASSWORD = "wtrack"
	POSTGRES_DB       = "PG_DATABASE"
	POSTGRES_SCHEMA   = "warning_tracker"
	POSTGRES_SSL_MODE = "disable"
	POSTGRES_PORT     = 5432
)

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	envArgs := []string{
		fmt.Sprintf("POSTGRES_USER=%s", POSTGRES_USER),
		fmt.Sprintf("POSTGRES_PASSWORD=%s", POSTGRES_PASSWORD),
		fmt.Sprintf("POSTGRES_DB=%s", POSTGRES_DB),
		fmt.Sprintf("POSTGRES_SCHEMA=%s", POSTGRES_SCHEMA),
		fmt.Sprintf("POSTGRES_SSL_MODE=%s", POSTGRES_SSL_MODE),
	}

	resource, err := pool.RunWithOptions(
		&dockertest.RunOptions{
			Name:       "wtrack_pg_test",
			Repository: "postgres",
			Tag:        "alpine",
			Env:        envArgs,
		})

	if resource == nil {
		fmt.Println(err)
		os.Exit(1)
	}

	IPAddress = resource.Container.NetworkSettings.IPAddress

	if err := pool.Retry(func() error {
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			resource.Container.NetworkSettings.IPAddress, POSTGRES_PORT, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_SSL_MODE)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)
		defer cancel()

		conn, err = sqlx.ConnectContext(ctx, "pgx", dsn)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		fmt.Println(resource.Container.NetworkSettings.IPAddress)
		log.Fatalf("Could not connect to database: %s", err)
	}

	defer func(conn *sqlx.DB) {
		if conn != nil {
			conn.Close()
		}
	}(conn)

	defer func(resource *dockertest.Resource) {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}(resource)

	err = conn.Ping()
	if err != nil {
		log.Printf("Could not ping to database: %s", err)
		return
	}

	migr, err := migrate.New(
		"file://../../../../db/migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", POSTGRES_USER, POSTGRES_PASSWORD, IPAddress, POSTGRES_PORT, POSTGRES_DB, POSTGRES_SSL_MODE),
	)

	if err != nil {
		log.Printf("failed create migrate instance %v", err)
		return
	}

	err = migr.Up()
	if err != nil {
		log.Printf("failed migrate up %v", err)
		return
	}

	code := m.Run()
	if code != 0 {
		log.Printf("failed test")
		return
	}

	err = migr.Down()
	if err != nil {
		log.Printf("failed migrate down %v", err)
		return
	}

}

func TestSomething(t *testing.T) {
	assert.Equal(t, nil, nil, "tmp test for save correct import assert")
}
