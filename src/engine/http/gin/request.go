package gin

import (
	"github.com/gin-gonic/gin"
	"io"
)

type GinRequest struct {
	context *gin.Context
}

func NewRequest(context *gin.Context) GinRequest {
	return GinRequest{context}
}

func (self GinRequest) Writer() io.Writer {
	return Flushed(self.context.Writer)
}

func (self GinRequest) Parameter(name string) string {
	return self.context.Params.ByName(name)
}
