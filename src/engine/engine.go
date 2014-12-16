package engine

import (
	"./command"
	"./config"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type Engine struct {
	configx config.Configuration
	router  *gin.Engine
}

type FlushedWriter struct {
	gin.ResponseWriter
}

func New() Engine {
	return Engine{config.New(), gin.Default()}
}

func (self FlushedWriter) Write(message []byte) (int, error) {
	wrote, err := self.ResponseWriter.Write(message)
	self.ResponseWriter.Flush()
	return wrote, err
}

func genericResponse(paramName, commandTemplate string) func(*gin.Context) {
	return func(context *gin.Context) {
		compiled := strings.Replace(commandTemplate, fmt.Sprintf("{%s}", paramName), context.Params.ByName(paramName), -1)
		command.ExecuteCommand(FlushedWriter{context.Writer}, compiled)
	}
}

func (self Engine) Draw() {
	for endpoint, commands := range self.configx.Commands {
		formattedEndpoint := fmt.Sprintf("/%s/:%s", endpoint, endpoint)

		for _, verbs := range commands {
			for verb, command := range verbs {
				handler := []gin.HandlerFunc{genericResponse(endpoint, command.(string))}
				self.router.Handle(strings.ToUpper(verb), formattedEndpoint, handler)
			}
		}
	}
}

func (self Engine) Start() {
	daemon := self.configx.Daemon
	listen := daemon.BindUrl()
	self.router.Run(listen)
}
