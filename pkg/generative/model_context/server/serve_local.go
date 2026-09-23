package server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) ServeLocal() {
	if e := server.ServeStdio(
		s.server,
		server.WithErrorLogger(log.NewGenericLogger()),
	); !errors.Canceled(e) {
		errors.PanicOnError(e)
	}
}
