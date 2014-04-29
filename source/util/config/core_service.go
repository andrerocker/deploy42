package config

import (
  "io"
  "fmt"
  "os/exec"
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

func executeCommand(output io.Writer, cmd string) {
  command := exec.Command("/usr/bin/bash", "-c", fmt.Sprintf("%s", cmd))
  command.Stdout = output
  command.Stderr = output
  command.Run()
}

func genericResponse(paramName, command string) func(martini.Params, http.ResponseWriter, *http.Request) {
  return func(params martini.Params, res http.ResponseWriter, req *http.Request) {
    compiled := strings.Replace(command, fmt.Sprintf("{%s}", paramName), params[paramName], -1)
    fmt.Println(compiled)
    res.WriteHeader(200)
    executeCommand(res, compiled)
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
