package gin

import (
	"github.com/andrerocker/deploy42/http"
	"github.com/gin-gonic/gin"
	"strings"
)

type GinAdapter struct {
	engine *gin.Engine
}

func New() GinAdapter {
	return GinAdapter{gin.Default()}
}

func (self GinAdapter) Start(bindUrl string) {
	self.engine.Run(bindUrl)
}

func (self GinAdapter) Register(method, endpoint string, handlers []http.Handler) {
	adapteds := self.buildHandlersList(handlers)
	self.engine.Handle(strings.ToUpper(method), endpoint, adapteds)
}

func (self GinAdapter) buildHandlersList(handlers []http.Handler) []gin.HandlerFunc {
	ginHandlers := make([]gin.HandlerFunc, 0)

	for _, handler := range handlers {
		ginHandlers = append(ginHandlers, self.buildHandler(handler))
	}

	return ginHandlers
}

func (self GinAdapter) buildHandler(handler http.Handler) func(context *gin.Context) {
	return func(context *gin.Context) {
		handler(NewRequest(context))
	}
}
