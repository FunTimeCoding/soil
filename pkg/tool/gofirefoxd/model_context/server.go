package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	firefox "github.com/funtimecoding/soil/pkg/tool/gofirefoxd/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	client   firefox.FirefoxSource
	reporter face.Reporter
}
