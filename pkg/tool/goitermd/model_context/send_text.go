package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SendText(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SendText,
) (*mcp.CallToolResult, error) {
	if e := s.service.SendText(
		a.SessionIdentifier,
		a.Text,
		a.SendEnter,
	); e != nil {
		return s.answer(e)
	}

	return response.Success("sent: %s", a.Text)
}
