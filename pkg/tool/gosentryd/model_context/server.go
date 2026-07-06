package model_context

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server       *server.MCPServer
	client       *sentry.Client
	organization string
	reporter     face.Reporter
}
