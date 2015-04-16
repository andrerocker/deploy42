package config

import "fmt"

type HttpFlags struct {
	Pipe bool
	Vars bool
}

type Daemon struct {
	Port int
	Bind string
	Pipe bool
	Load []string
	Http HttpFlags
}

func (self Daemon) BindUrl() string {
	return fmt.Sprintf("%s:%d", self.Bind, self.Port)
}
