package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	forge "github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	client   forge.Forge
	reporter face.Reporter
}
