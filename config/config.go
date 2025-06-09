package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Configuration represents the application configuration structure.
type Configuration struct {
	API struct {
		Key      string `yaml:"key"`
		Endpoint string `yaml:"endpoint"`
	} `yaml:"api"`
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
}

func LoadConfig(path string) Configuration {
	var config Configuration

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error unmarshalling yaml: %v", err)
	}

	return config
}
