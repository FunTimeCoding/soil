package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	iterm "github.com/funtimecoding/soil/pkg/tool/goitermd/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	client   iterm.ItermSource
	reporter face.Reporter
}
