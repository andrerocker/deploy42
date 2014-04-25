package config

import "fmt"

type CoreService struct {
  Port int
  Bind string
}

type Command struct {
  Pattern string
  Verb map[string]string
}

func (self CoreService) BindUrl() string {
  return fmt.Sprintf("%s:%d", self.Bind, self.Port)
}
