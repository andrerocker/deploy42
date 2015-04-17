package deploy42

import (
	"fmt"
	"github.com/andrerocker/deploy42/command"
	"github.com/andrerocker/deploy42/config"
	"github.com/andrerocker/deploy42/http"
	"github.com/andrerocker/deploy42/http/gin"
	"io"
	"strings"
)

type Engine struct {
	http    http.Engine
	config  config.Configuration
	filters map[string]http.Handler
}

func New(configFile string) Engine {
	return Engine{gin.New(), config.New(configFile), make(map[string]http.Handler)}
}

func (self Engine) RegisterFilter(name string, filter http.Handler) {
	self.filters[name] = filter
}

func (self Engine) Draw() {
	for _, namespace := range self.config.Namespaces {
		for groupName, commands := range namespace.Commands {
			for _, verbs := range commands {
				for verb, command := range verbs {
					route := self.formattedEndpoint(namespace, groupName)
					handlers := self.wrapValuesHandler(namespace, groupName, command.(string))

					self.http.Register(verb, route, handlers)
				}
			}
		}
	}
}

func (self Engine) Start() {
	daemon := self.config.Daemon
	listen := daemon.BindUrl()

	self.http.Start(listen)
}

func (self Engine) wrapValuesHandler(namespace config.Namespace, groupName, commandTemplate string) []http.Handler {
	filters := self.resolveNamespaceFilters(namespace)
	return append(filters, func(request http.Request) {
		reader := self.resolveReader(request)
		target := fmt.Sprintf("{%s}", groupName)
		content := request.ContextParameter(groupName)
		compiled := strings.Replace(commandTemplate, target, content, -1)

		command.ExecuteCommand(reader, request.Writer(), compiled)
	})
}

func (self Engine) formattedEndpoint(namespace config.Namespace, groupName string) string {
	if self.config.Daemon.Http.Vars {
		return fmt.Sprintf("/%s/%s/*%s", namespace.Endpoint, groupName, groupName)
	}

	return fmt.Sprintf("/%s/%s", namespace.Endpoint, groupName)
}

func (self Engine) resolveReader(request http.Request) io.Reader {
	if self.config.Daemon.Http.Pipe {
		return request.Reader()
	}

	return strings.NewReader("")
}

func (self Engine) resolveNamespaceFilters(namespace config.Namespace) []http.Handler {
	filters := make([]http.Handler, 0)
	for _, filterName := range namespace.Chaining {
		currentFilter := self.filters[filterName]
		if currentFilter != nil {
			filters = append(filters, currentFilter)
		}
	}

	return filters
}
