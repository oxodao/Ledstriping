package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/oxodao/ledstrip/models"
)

type Config struct {
	ListeningUrl string
	SerialPort   string

	Favorites []models.Favorite
}

func getConfigPath() string {
	if _, err := os.Stat("config.json"); err == nil {
		return "config.json"
	}

	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".config", "ledstrip")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}

	path = filepath.Join(path, "config.json")
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return path
	}

	panic("no config file found")
}

func Load(path string) (*Config, error) {
	if len(path) == 0 {
		path = getConfigPath()
	}

	data, err := ioutil.ReadFile(path)
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
