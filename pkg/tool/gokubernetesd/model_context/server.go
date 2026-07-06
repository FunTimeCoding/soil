package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/mark3labs/mcp-go/server"
	"sync"
)

type Server struct {
	server   *server.MCPServer
	service  *service.Service
	sessions sync.Map
	readOnly bool
	reporter face.Reporter
}
