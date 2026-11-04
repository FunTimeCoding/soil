package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listClient(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.Empty,
) (*mcp.CallToolResult, error) {
	rows, e := s.service.ListClients()

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(convert.Clients(rows))
}
