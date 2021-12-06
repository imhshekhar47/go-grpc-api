package core

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type serverConfig struct {
	Mode string `yaml:"mode"`
	Port int16  `yaml:"port"`
}

type applicationConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

//AppConfig definition
//Description: main configuration model
type AppConfig struct {
	Server      serverConfig      `yaml:"server"`
	Application applicationConfig `yaml:"application"`
}

var instance *AppConfig

//GetAppConfig return config from application.yaml
func GetAppConfig() AppConfig {
	if nil == instance {
		config := AppConfig{
			Server: serverConfig{
				Port: 8080,
			},
			Application: applicationConfig{
				Name:    "go-grpc-api",
				Version: "v1-SNAPSHOT",
			},
		}

		yamlfile, err := ioutil.ReadFile("application.yaml")

		if err != nil {
			log.Printf("Error while reading application.yaml: %v\n", err)
			return config
		}

		err = yaml.Unmarshal(yamlfile, &config)

		if err != nil {
			log.Printf("Error while loading application.yml: %v\n", err)
			return config
		}
		instance = &config
	}

	return *instance
}
