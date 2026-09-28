package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Play(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Play,
) (*mcp.CallToolResult, error) {
	if a.SessionIdentifier == "" {
		return response.Fail("session_id is required")
	}

	if len(a.ItemIDs) == 0 {
		return response.Fail("item_ids is required")
	}

	if a.PlayCommand == "" {
		a.PlayCommand = "PlayNow"
	}

	e := s.client.Play(
		a.SessionIdentifier,
		a.ItemIDs,
		a.PlayCommand,
		a.StartPositionTicks,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success("playback started")
}
