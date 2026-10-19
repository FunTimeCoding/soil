package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server    *server.MCPServer
	service   *service.Service
	selection *selection.Selection
	reporter  face.Reporter
}
