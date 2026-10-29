package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ReadHistory(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ReadHistory,
) (*mcp.CallToolResult, error) {
	v, e := s.service.ReadHistory(a.SessionIdentifier, int(a.Count))

	if e != nil {
		return s.answer(e)
	}

	return response.SuccessAny(v)
}
