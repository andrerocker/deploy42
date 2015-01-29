package config

import (
	"github.com/imdario/mergo"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"path/filepath"
)

func YAMLoad(configFile string) Configuration {
	cfg := load(configFile)
	extendPath := cfg.Daemon.Load

	if extendPath != "" {
		files, _ := filepath.Glob(extendPath)

		for _, file := range files {
			internal := load(file).Commands
			mergo.Merge(&cfg.Commands, internal)
		}
	}

	return cfg
}

func load(configFile string) Configuration {
	cfg := Configuration{}
	data, _ := ioutil.ReadFile(configFile)
	yaml.Unmarshal(data, &cfg)
	return cfg
}
