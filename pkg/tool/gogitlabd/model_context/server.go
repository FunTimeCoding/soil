package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/mark3labs/mcp-go/server"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type Server struct {
	server   *server.MCPServer
	client   *gitlab.Client
	reporter face.Reporter
}
