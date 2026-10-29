package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	provision "github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	service  provision.Provision
	reporter face.Reporter
}
