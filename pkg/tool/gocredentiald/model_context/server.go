package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	credentialFace "github.com/funtimecoding/soil/pkg/tool/gocredentiald/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	service  credentialFace.CredentialSource
	reporter face.Reporter
}
