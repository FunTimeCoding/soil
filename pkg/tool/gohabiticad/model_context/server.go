package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	habitica "github.com/funtimecoding/soil/pkg/tool/gohabiticad/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	habitica habitica.HabiticaSource
	reporter face.Reporter
}
