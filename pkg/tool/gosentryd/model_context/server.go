package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/reacher"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/service"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	service  *service.Service
	host     string
	reporter face.Reporter
	reacher  *reacher.Reacher
}
