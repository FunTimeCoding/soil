package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	opnsense "github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	opnsense opnsense.OpnsenseSource
	reporter face.Reporter
}
