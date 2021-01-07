package store

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
