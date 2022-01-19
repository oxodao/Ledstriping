package main

import (
	"flag"

	"github.com/oxodao/ledstrip/config"
	"github.com/oxodao/ledstrip/services"
)

func main() {
	configPath := flag.String("config", "", "Path to config file")

	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	prv, err := services.NewProvider(cfg)
	if err != nil {
		panic(err)
	}

	RunServer(prv)
}
