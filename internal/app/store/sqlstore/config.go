package sqlstore

import "fmt"

type Config struct {
	Host     string `toml:"host"`
	DBName   string `toml:"db_name"`
	DBPort   int    `toml:"db_port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	SSLMode  string `toml:"ssl_mode"`
}

func NewConfig() *Config {
	return &Config{
		Host:     "db",
		DBName:   "restapi",
		DBPort:   5432,
		User:     "postgres",
		Password: "mysecretpassword",
		SSLMode:  "disable",
	}
}

func (c *Config) DBUrl() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.DBPort,
		c.User,
		c.Password,
		c.DBName,
		c.SSLMode,
	)
}
