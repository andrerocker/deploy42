package config

import (
	"gopkg.in/yaml.v1"
	"io/ioutil"
)

func YAMLoad(configFile string) Configuration {
	cfg := Configuration{}
	data, _ := ioutil.ReadFile(configFile)
	yaml.Unmarshal(data, &cfg)
	return cfg
}
