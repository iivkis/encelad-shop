package config

import (
	"fmt"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
)

type conf struct {
	PostgresDB struct {
		Name     string `env:"postgres_database_name" validate:"required"`
		Host     string `env:"postgres_database_host" validate:"required"`
		Port     int    `env:"postgres_database_port" validate:"required"`
		User     string `env:"postgres_database_user" validate:"required"`
		Password string `env:"postgres_database_password" validate:"required"`
		SSLMode  string `env:"postgres_database_sslmode" validate:"required"`
	}

	Server struct {
		Host string `env:"server_host" validate:"required"`
		Port int    `env:"server_port" validate:"required"`
	}
}

func (c *conf) ServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}

func (c *conf) PostgresConnStr() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.PostgresDB.User,
		c.PostgresDB.Password,
		c.PostgresDB.Host,
		c.PostgresDB.Port,
		c.PostgresDB.Name,
		c.PostgresDB.SSLMode,
	)
}

func (c *conf) Validate() error {
	return validator.New().Struct(c)
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

		if err := config.Validate(); err != nil {
			panic(err)
		}
	})
	return &config
}
