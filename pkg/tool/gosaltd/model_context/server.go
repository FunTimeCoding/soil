package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	salt "github.com/funtimecoding/soil/pkg/tool/gosaltd/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	service  salt.Salt
	reporter face.Reporter
}
