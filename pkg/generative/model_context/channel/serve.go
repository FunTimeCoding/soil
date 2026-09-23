package channel

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) Serve() {
	if e := server.ServeStdio(s.server); !errors.Canceled(e) {
		errors.PanicOnError(e)
	}
}
