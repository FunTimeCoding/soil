package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	client   *netbox.Client
	store    *store.Store
	reporter face.Reporter
}
