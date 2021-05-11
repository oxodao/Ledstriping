package main

import (
	"fmt"
	"github.com/oxodao/ledstrip/config"
	"github.com/oxodao/ledstrip/services"
)

func main() {
	fmt.Println("Ledstrip")

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	prv, err := services.NewProvider(cfg)
	if err != nil {
		panic(err)
	}

	RunServer(prv)
}