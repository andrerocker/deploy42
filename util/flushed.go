package util

import "github.com/gin-gonic/gin"

type FlushedWriter struct {
	gin.ResponseWriter
}

func Flushed(writer gin.ResponseWriter) FlushedWriter {
	return FlushedWriter{writer}
}

func (self FlushedWriter) Write(message []byte) (int, error) {
	wrote, err := self.ResponseWriter.Write(message)
	self.ResponseWriter.Flush()
	return wrote, err
}
