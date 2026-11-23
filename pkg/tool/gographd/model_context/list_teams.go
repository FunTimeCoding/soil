package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gographd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListTeams(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListTeams,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Teams()

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(result)
}
