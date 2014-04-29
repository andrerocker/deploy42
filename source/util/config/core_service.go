package config

import (
  "fmt"
  "strings"
  "net/http"
  "github.com/andrerocker/martini"
)

type Command map[string]interface{}
type CommandList map[string][]Command

type CoreService struct {
  Port int
  Bind string
}

type Configuration struct {
  Service Command
  Commands CommandList
}

func (self CoreService) BindUrl() string {
  return fmt.Sprintf("%s:%d", self.Bind, self.Port)
}

func genericResponse(paramName, command string) func(martini.Params, http.ResponseWriter, *http.Request) {
  return func(params martini.Params, res http.ResponseWriter, req *http.Request) {
    compiled := fmt.Sprintf("%s - %s", command, params[paramName])
    res.WriteHeader(200)
    fmt.Fprintf(res, compiled)
  }
}

func (self CommandList) DrawRoutes(router *martini.ClassicMartini) {
  for endpoint, commands := range self {
    formattedEndpoint := fmt.Sprintf("/%s/:%s", endpoint, endpoint)

    for _, verbs := range commands {
      for verb, command := range verbs {
        router.Add(strings.ToUpper(verb), formattedEndpoint, genericResponse(endpoint, command.(string)))
      }
    }
  }
}
