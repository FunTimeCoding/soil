package channel

import (
	"github.com/funtimecoding/soil/pkg/errors"
	generative "github.com/funtimecoding/soil/pkg/generative/model_context/server"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) Serve() {
	if e := server.ServeStdio(
		s.server,
		server.WithStdioContextFunc(generative.LegacyProtocol),
	); !errors.Canceled(e) {
		errors.PanicOnError(e)
	}
}
