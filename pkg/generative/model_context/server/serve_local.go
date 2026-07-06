package server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) ServeLocal() {
	errors.PanicOnError(
		server.ServeStdio(
			s.server,
			server.WithErrorLogger(log.NewGenericLogger()),
		),
	)
}
