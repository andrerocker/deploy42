package config

type Configuration struct {
	Daemon   Daemon
	Commands CommandList
}

func New(configFile string) Configuration {
	return YAMLoad(configFile)
}
