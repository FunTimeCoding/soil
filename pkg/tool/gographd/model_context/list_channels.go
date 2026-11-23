package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gographd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListChannels(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListChannels,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Channels(a.Team)

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(result)
}
