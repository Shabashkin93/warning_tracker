package config

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug  bool   `env:"INFO" env-default:"false" env-upd:"true"`
	LogLevel string `env:"LOG_LEVEL" env-default:"info" env-upd:"true"`
	DB       struct {
		User     string `env:"POSTGRES_USER" env-default:"wtrack"`
		Password string `env:"POSTGRES_PASSWORD" env-default:"wtrack"`
		Db       string `env:"POSTGRES_DB" env-default:"PG_DATABASE"`
		Port     string `env:"POSTGRES_PORT" env-default:"5432"`
		SslMode  string `env:"POSTGRES_SSL_MODE" env-default:"disable"`
		Address  string `env:"POSTGRES_ADDRESS" env-default:"0.0.0.0"`
		Schema   string `env:"POSTGRES_SCHEMA" env-default:"warning_tracker"`
		Timeout  int    `env:"POSTGRES_TIMEOUT" env-default:"2"`
		Table    struct {
			Warning string `env:"DB_TB_WARNING" env-default:"warning"`
		}
	}
	HTTP struct {
		Port string `env:"SERVER_PORT" env-default:"8090"`
		URL  struct {
			Warning string `env:"URL_WARNING" env-default:"/warning"`
			Status  string `env:"URL_STATUS" env-default:"/status"`
		}
	}
	REDIS struct {
		Address  string `env:"REDIS_ADDRESS" env-default:"0.0.0.0"`
		Port     string `env:"REDIS_PORT" env-default:"6379"`
		Password string `env:"REDIS_PASSWORD" env-default:"admin"`
		Timeout  int    `env:"REDIS_TIMEOUT" env-default:"2"`
	}
}

var instance *Config

var once sync.Once

func GetConfig() (config *Config, err error) {
	once.Do(func() {
		instance = &Config{}
		if err = cleanenv.ReadEnv(instance); err != nil {
			help, errDescr := cleanenv.GetDescription(instance, nil)
			if errDescr == nil {
				fmt.Println(help)
			}
		}
	})
	config = instance
	return
}
