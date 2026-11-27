package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gographd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SearchUsers(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.SearchUsers,
) (*mcp.CallToolResult, error) {
	result, e := s.service.SearchUsers(a.Query, a.Limit)

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(result)
}
