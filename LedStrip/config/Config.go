package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/oxodao/ledstrip/models"
)

type Config struct {
	ListeningUrl string
	SerialPort   string

	Favorites []models.Favorite
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

func (c Config) Save() error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile("config.json", data, 0644)
}
