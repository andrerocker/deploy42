package config

import "fmt"

type Daemon struct {
	Port int
	Bind string
}

func (self Daemon) BindUrl() string {
	return fmt.Sprintf("%s:%d", self.Bind, self.Port)
}
