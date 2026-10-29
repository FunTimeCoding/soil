package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	secret "github.com/funtimecoding/soil/pkg/tool/gosecretd/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	service  secret.Secrets
	reporter face.Reporter
}
