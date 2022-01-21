package services

import (
	"github.com/oxodao/ledstrip/config"
	"github.com/oxodao/ledstrip/sdk"
)

type Provider struct {
	Config   *config.Config
	Ledstrip *sdk.Ledstrip
}

func NewProvider(cfg *config.Config) (*Provider, error) {
	sdk, err := sdk.Connect(cfg.SerialPort)
	if err != nil {
		panic(err)
	}

	prv := &Provider{
		Config:   cfg,
		Ledstrip: sdk,
	}

	return prv, nil
}
