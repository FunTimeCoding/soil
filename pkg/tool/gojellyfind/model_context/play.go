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
	e := s.service.Play(
		a.SessionIdentifier,
		a.ItemIDs,
		a.PlayCommand,
		a.StartPositionTicks,
	)

	if e != nil {
		return s.answer(e)
	}

	return response.Success("playback started")
}
