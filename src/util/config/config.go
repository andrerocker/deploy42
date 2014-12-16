package config

import (
	"gopkg.in/yaml.v1"
	"io/ioutil"
)

var config Configuration

func unmarshal(configFile string) Configuration {
	deploy := Configuration{}
	data, _ := ioutil.ReadFile(configFile)

	yaml.Unmarshal(data, &deploy)
	return deploy
}

func NewCoreService(configFile string) CoreService {
	config := unmarshal(configFile)
	service := config.Service
	return CoreService{service["port"].(int), service["bind"].(string)}
}

func NewServiceCommands() CommandList {
	return config.Commands
}
