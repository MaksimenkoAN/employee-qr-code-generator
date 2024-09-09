package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

var AppConfig Config

func LoadConfig(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		return err
	}

	return nil
}
