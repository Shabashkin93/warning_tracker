package redis

import (
	"context"
	"time"

	"github.com/Shabashkin93/warning_tracker/internal/logging"
	"github.com/Shabashkin93/warning_tracker/internal/project_errors"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Client  *redis.Client
	ctx     context.Context
	timeout time.Duration
}

func Init(ctx context.Context, logger logging.Logger, address, port, password string, timeout time.Duration) (Cache, error) {
	client := Cache{}
	client.Client = redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	client.ctx = ctx
	client.timeout = timeout

	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()

	_, err := client.Client.Ping(ctx).Result()
	if err != nil {
		logger.Info(ctx, "Cache connection established")
	}

	return client, err
}

func (cache Cache) Shutdown() (err error) {
	err = cache.Client.Close()
	return
}

func (c Cache) Set(key string, value interface{}) (err error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.timeout*time.Second)
	defer cancel()

	err = c.Client.Set(ctx, key, value, 0).Err()
	return
}

func (c Cache) Get(key string) (value string, err error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.timeout*time.Second)
	defer cancel()

	value, err = c.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		err = project_errors.CacheKeyNotFound
	}
	return
}

func (c Cache) Delete(key string) (err error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.timeout*time.Second)
	defer cancel()

	_, err = c.Client.Del(ctx, key).Result()
	return
}
