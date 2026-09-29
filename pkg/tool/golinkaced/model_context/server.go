package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	linkace "github.com/funtimecoding/soil/pkg/tool/golinkaced/face"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/service"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	client   linkace.LinkAceSource
	service  *service.Service
	reporter face.Reporter
}
