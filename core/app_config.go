package core

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

//ServerConfig definition
type ServerConfig struct {
	Mode string `yaml:"mode"`
	Port int16  `yaml:"port"`
}

//AppConfig definition
type AppConfig struct {
	Server ServerConfig `yaml:"server"`
}

//GetAppConfig return config from pplication.yaml
func GetAppConfig() AppConfig {
	config := AppConfig{
		Server: ServerConfig{
			Port: 8080,
		},
	}

	yamlfile, err := ioutil.ReadFile("application.yml")

	if err != nil {
		log.Printf("Error while reading application.yaml: %v\n", err)
		return config
	}

	err = yaml.Unmarshal(yamlfile, &config)

	if err != nil {
		log.Printf("Error while loading application.yml: %v\n", err)
		return config
	}

	return config
}
