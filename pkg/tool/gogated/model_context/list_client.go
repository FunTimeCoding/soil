package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listClient(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.Empty,
) (*mcp.CallToolResult, error) {
	return response.SuccessAny(convert.Clients(s.service.ListClients()))
}
