package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/backup"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/catalog"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/query"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server    *server.MCPServer
	catalog   *catalog.Service
	sql       *query.Service
	backup    *backup.Service
	instances *instance_tool.Tool[instance.Instance, *instance.Instance]
	reporter  face.Reporter
}
