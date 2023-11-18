package config

import (
	"errors"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port         string        `yaml:"port"`
	Host         string        `yaml:"host"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

func Load(path string) (*Config, error) {
	configInfo, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("error when trying to read a httpConfig file or load data from this file" + err.Error())
	}

	var cfg Config
	err = yaml.Unmarshal(configInfo, cfg)
	if err != nil {
		return nil, errors.New("Failed to unmarshal date from httpConfig file: " + err.Error())
	}

	return &cfg, nil

}
