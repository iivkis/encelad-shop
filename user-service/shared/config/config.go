package config

import (
	"sync"

	"github.com/caarlos0/env/v11"
)

type conf struct {
	PostgresDB struct {
		Name     string `env:"postgres_database_name"`
		Host     string `env:"postgres_database_host"`
		Port     int    `env:"postgres_database_port"`
		User     string `env:"postgres_database_user"`
		Password string `env:"postgres_database_password"`
	}
}

var (
	onceConfig sync.Once
	config     conf = conf{}
)

func Config() *conf {
	onceConfig.Do(func() {
		if err := env.Parse(&config); err != nil {
			panic(err)
		}
	})
	return &config
}
