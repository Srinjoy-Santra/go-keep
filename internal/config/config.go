package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func NewConfig(env string) (*Configuration, error) {
	config, err := loadConfig(env)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func loadConfig(env string) (*Configuration, error) {

	configPath := filepath.Join("internal", "config", "tier", fmt.Sprintf("%s.yml", env))
	conf := Configuration{}
	err := decodeYml(configPath, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func decodeYml(filePath string, config interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Could not open file")
		return err
	}

	decoder := yaml.NewDecoder(file)
	defer file.Close()
	err = decoder.Decode(config)
	if err != nil {
		log.Println("Could not decode file")
		return err
	}
	return nil
}
