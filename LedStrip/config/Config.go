package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ListeningUrl string
	SerialPort string
}

func Load() (*Config, error) {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = json.Unmarshal(data, cfg)

	return cfg, err
}