package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	sentry "github.com/funtimecoding/soil/pkg/tool/gosentryd/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server       *server.MCPServer
	client       sentry.SentrySource
	organization string
	reporter     face.Reporter
}
