package deploy42

import (
	"fmt"
	"github.com/andrerocker/deploy42/command"
	"github.com/andrerocker/deploy42/config"
	"github.com/andrerocker/deploy42/util"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
)

type Engine struct {
	http    *gin.Engine
	config  config.Configuration
	filters map[string]gin.HandlerFunc
}

func New(configFile string) Engine {
	return Engine{gin.Default(), config.New(configFile), make(map[string]gin.HandlerFunc)}
}

func (self Engine) Chaining(name string, filter gin.HandlerFunc) {
	self.filters[name] = filter
}

func (self Engine) Draw() {
	for _, namespace := range self.config.Namespaces {
		for groupName, commands := range namespace.Commands {
			for _, verbs := range commands {
				for verb, command := range verbs {
					route := self.formattedEndpoint(namespace, groupName)
					handlers := self.wrapValuesHandler(namespace, groupName, command.(string))

					self.http.Handle(strings.ToUpper(verb), route, handlers)
				}
			}
		}
	}
}

func (self Engine) Start() {
	daemon := self.config.Daemon
	listen := daemon.BindUrl()

	self.http.Run(listen)
}

func (self Engine) wrapValuesHandler(namespace config.Namespace, groupName, commandTemplate string) []gin.HandlerFunc {
	filters := self.resolveNamespaceFilters(namespace)
	return append(filters, func(context *gin.Context) {
		reader := self.resolveReader(context)
		target := fmt.Sprintf("{%s}", groupName)
		content := context.Params.ByName(groupName)
		compiled := strings.Replace(commandTemplate, target, content, -1)

		command.ExecuteCommand(reader, util.Flushed(context.Writer), compiled)
	})
}

func (self Engine) formattedEndpoint(namespace config.Namespace, groupName string) string {
	if self.config.Daemon.Http.Vars {
		return fmt.Sprintf("/%s/%s/*%s", namespace.Endpoint, groupName, groupName)
	}

	return fmt.Sprintf("/%s/%s", namespace.Endpoint, groupName)
}

func (self Engine) resolveReader(context *gin.Context) io.Reader {
	if self.config.Daemon.Http.Pipe {
		return context.Request.Body
	}

	return strings.NewReader("")
}

func (self Engine) resolveNamespaceFilters(namespace config.Namespace) []gin.HandlerFunc {
	filters := make([]gin.HandlerFunc, 0)
	for _, filterName := range namespace.Chaining {
		currentFilter := self.filters[filterName]
		if currentFilter != nil {
			filters = append(filters, currentFilter)
		}
	}

	return filters
}
