package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) PlaybackCommand(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.PlaybackCommand,
) (*mcp.CallToolResult, error) {
	if a.SessionIdentifier == "" {
		return response.Fail("session_id is required")
	}

	if a.Command == "" {
		return response.Fail("command is required")
	}

	e := s.client.PlaybackCommand(
		a.SessionIdentifier,
		a.Command,
		a.SeekPositionTicks,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success(fmt.Sprintf("%s sent", a.Command))
}
