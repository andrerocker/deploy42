package config

type Configuration struct {
	Daemon     Daemon
	Namespaces NamespaceList
}

func New(configFile string) Configuration {
	return YAMLoad(configFile)
}
