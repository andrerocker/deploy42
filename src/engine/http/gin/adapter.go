package gin

import (
	"../../http"
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

func (self GinAdapter) Register(method, endpoint string, handler http.Handler) {
	adapted := self.buildHandler(handler)
	self.engine.Handle(strings.ToUpper(method), endpoint, []gin.HandlerFunc{adapted})
}

func (self GinAdapter) buildHandler(handler http.Handler) func(context *gin.Context) {
	return func(context *gin.Context) {
		handler(NewRequest(context))
	}
}
