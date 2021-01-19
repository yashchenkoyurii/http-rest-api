package apiserver

import (
	"github.com/yashchenkoyurii/http-rest-api/internal/app/store/sqlstore"
)

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Sql      *sqlstore.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8888",
		LogLevel: "debug",
		Sql:      sqlstore.NewConfig(),
	}
}
