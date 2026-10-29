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
	e := s.service.SetVolume(a.SessionIdentifier, a.Level)

	if e != nil {
		return s.answer(e)
	}

	return response.Success(fmt.Sprintf("volume set to %d", a.Level))
}
