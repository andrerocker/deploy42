package http

import "io"

type Request interface {
	Writer() io.Writer
	Parameter(string) string
}
