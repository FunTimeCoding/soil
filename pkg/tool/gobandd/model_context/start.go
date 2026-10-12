package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) start(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.Start,
) (*mcp.CallToolResult, error) {
	return s.power(x, constant.OperationStart, false)
}
