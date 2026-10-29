package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) UnsubscribeThread(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.UnsubscribeThread,
) (*mcp.CallToolResult, error) {
	if a.Root == "" {
		return response.Fail("root is required")
	}

	if a.Callsign == "" {
		return response.Fail("callsign is required")
	}

	removed, e := s.service.Unsubscribe(a.Callsign, a.Root)

	if e != nil {
		return s.answer(e)
	}

	if !removed {
		return response.Success("no such subscription")
	}

	return response.Success("unsubscribed")
}
