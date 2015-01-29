package config

import (
	"github.com/imdario/mergo"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"path/filepath"
)

func YAMLoad(configFile string) Configuration {
	cfg := loadBase(configFile)
	mergo.Merge(&cfg.Commands, loadExtensionList(cfg.Daemon.Load))

	return cfg
}

func loadExtensionList(configFiles []string) CommandList {
	commandList := make(CommandList)

	if configFiles != nil {
		for _, currentConfigFile := range configFiles {
			mergo.Merge(&commandList, loadExtensionGlob(currentConfigFile))
		}
	}

	return commandList
}

func loadExtensionGlob(configFile string) CommandList {
	files, _ := filepath.Glob(configFile)
	commandList := make(CommandList)

	for _, file := range files {
		internal := loadBase(file).Commands
		mergo.Merge(&commandList, internal)
	}

	return commandList
}

func loadBase(configFile string) Configuration {
	cfg := Configuration{}
	data, _ := ioutil.ReadFile(configFile)
	yaml.Unmarshal(data, &cfg)
	return cfg
}
