package http

import "io"

type Handler func(Request)

type Engine interface {
	Start(string)
	Register(string, string, Handler)
}

type Request interface {
	Writer() io.Writer
	Parameter(string) string
}
