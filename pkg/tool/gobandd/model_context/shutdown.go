package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) shutdown(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Shutdown,
) (*mcp.CallToolResult, error) {
	return s.power(x, constant.OperationShutdown, a.Force)
}
