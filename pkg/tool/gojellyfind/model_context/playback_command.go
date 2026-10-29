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
	e := s.service.PlaybackCommand(
		a.SessionIdentifier,
		a.Command,
		a.SeekPositionTicks,
	)

	if e != nil {
		return s.answer(e)
	}

	return response.Success(fmt.Sprintf("%s sent", a.Command))
}
