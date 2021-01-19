package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"

	"github.com/yashchenkoyurii/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "/go/bin/configs/apiserver.toml", "Path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	app := apiserver.NewApp(config)

	if err := app.Run(); err != nil {
		log.Fatal()
	}
}
