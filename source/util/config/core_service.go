package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os/exec"
	"strings"
)

type Command map[string]interface{}
type CommandList map[string][]Command

type CoreService struct {
	Port int
	Bind string
}

type Configuration struct {
	Service  Command
	Commands CommandList
}

func (self CoreService) BindUrl() string {
	return fmt.Sprintf("%s:%d", self.Bind, self.Port)
}

type FlushedWriter struct {
	gin.ResponseWriter
}

func (self FlushedWriter) Write(message []byte) (int, error) {
	wrote, err := self.ResponseWriter.Write(message)
	self.ResponseWriter.Flush()
	return wrote, err
}

func executeCommand(output FlushedWriter, cmd string) {
	command := exec.Command("/bin/bash", "-c", fmt.Sprintf("%s", cmd))
	command.Stdout = output
	command.Stderr = output
	command.Run()
}

func genericResponse(paramName, command string) func(*gin.Context) {
	return func(context *gin.Context) {
		compiled := strings.Replace(command, fmt.Sprintf("{%s}", paramName), context.Params.ByName(paramName), -1)
		fmt.Println(compiled)
		executeCommand(FlushedWriter{context.Writer}, compiled)
	}
}

func (self CommandList) DrawRoutes(router *gin.Engine) {
	for endpoint, commands := range self {
		formattedEndpoint := fmt.Sprintf("/%s/:%s", endpoint, endpoint)

		for _, verbs := range commands {
			for verb, command := range verbs {
				router.Handle(strings.ToUpper(verb), formattedEndpoint, []gin.HandlerFunc{genericResponse(endpoint, command.(string))})
			}
		}
	}
}
