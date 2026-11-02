package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/file"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/packages"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/pattern"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/reference"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server    *server.MCPServer
	file      *file.Service
	module    *module.Service
	packages  *packages.Service
	pattern   *pattern.Service
	reference *reference.Service
	symbol    *symbol.Service
	selection *selection.Selection
	reporter  face.Reporter
}
