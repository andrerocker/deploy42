package config

import (
	"flag"
)

type Configuration struct {
	Daemon   Daemon
	Commands CommandList
}

func New() Configuration {
	return YAMLoad(parseCommandLineArgs())
}

func parseCommandLineArgs() string {
	configFile := flag.String("config", "deploy42.yml", "configuration file")
	flag.Parse()
	return *configFile
}
