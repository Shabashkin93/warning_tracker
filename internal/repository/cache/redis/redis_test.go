package redis_test

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/Shabashkin93/warning_tracker/internal/config"
	"github.com/Shabashkin93/warning_tracker/internal/repository"
	"github.com/Shabashkin93/warning_tracker/internal/repository/cache/redis"
	"github.com/Shabashkin93/warning_tracker/pkg/logging"
	logger "github.com/Shabashkin93/warning_tracker/pkg/logging/slog"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

var IPAddress string
var repos *repository.Repository
var cache redis.Cache

const (
	REDIS_PASSWORD    = "wtrack"
	REDIS_PORT        = "6379"
	REDIS_BACKUP_TIME = 2
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

	var ctx = context.Background()

	cfg, err := config.GetConfig()
	if err != nil {
		slog.Error("Failed get config")
		os.Exit(1)
	}

	out := os.Stdout
	loggerEnt := logger.NewLogger(cfg.LogLevel, out)
	defer out.Close()

	logger := logging.NewLogger(loggerEnt)

	envArgs := []string{
		fmt.Sprintf("REDIS_PASSWORD=%s", REDIS_PASSWORD),
		fmt.Sprintf("REDIS_PORT=%s", REDIS_PORT),
	}

	resource, err := pool.RunWithOptions(
		&dockertest.RunOptions{
			Name:       "wtrack_redis_test",
			Repository: "redis",
			Tag:        "alpine",
			Env:        envArgs,
		})

	if resource == nil {
		fmt.Println(err)
		os.Exit(1)
	}

	IPAddress = resource.Container.NetworkSettings.IPAddress

	if err := pool.Retry(func() error {
		cache, err = redis.Init(ctx, logger, resource.Container.NetworkSettings.IPAddress, REDIS_PORT, REDIS_PASSWORD, time.Duration(REDIS_BACKUP_TIME))
		if err != nil {
			logger.Fatal(ctx, "Could not set up cache", slog.String("error", fmt.Sprintf("%v", err)))
		}

		repos = repository.NewRepository(ctx, logger, cfg, nil, cache)
		return nil
	}); err != nil {
		fmt.Println(resource.Container.NetworkSettings.IPAddress)
		log.Fatalf("Could not connect to redis: %s", err)
	}

	defer func(cache *redis.Cache) {
		if cache != nil {
			cache.Shutdown()
		}
	}(&cache)

	defer func(resource *dockertest.Resource) {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}(resource)

	defer repos.Stop()

	code := m.Run()
	if code != 0 {
		log.Printf("failed test")
		return
	}

}

func TestWarning(t *testing.T) {
	err := repos.Cache.Set("key1", "value1")
	assert.Equal(t, err, nil, "failed create record in redis err:%v", err)
	value, err := repos.Cache.Get("key1")
	assert.Equal(t, err, nil, "failed get record in redis err:%v", err)
	assert.Equal(t, value, "value1", "failed get value record in redis")
}
