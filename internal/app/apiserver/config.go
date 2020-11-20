package apiserver

type Config struct {
	bindAddr string `toml:"bind_addr"`
	logLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		bindAddr: ":8888",
		logLevel: "debug",
	}
}
