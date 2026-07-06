package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	sublime "github.com/funtimecoding/soil/pkg/tool/gosublimed/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	client   sublime.SublimeSource
	reporter face.Reporter
}
