package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SetVolume(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SetVolume,
) (*mcp.CallToolResult, error) {
	if a.SessionIdentifier == "" {
		return response.Fail("session_id is required")
	}

	if a.Level < 0 || a.Level > 100 {
		return response.Fail("level must be between 0 and 100")
	}

	e := s.client.SetVolume(a.SessionIdentifier, a.Level)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success(fmt.Sprintf("volume set to %d", a.Level))
}
