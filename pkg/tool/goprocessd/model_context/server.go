package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	supervisor "github.com/funtimecoding/soil/pkg/tool/goprocessd/server"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server     *server.MCPServer
	supervisor *supervisor.Server
	reporter   face.Reporter
}
