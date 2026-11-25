package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gographd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SendChatMessage(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.SendChatMessage,
) (*mcp.CallToolResult, error) {
	result, e := s.service.SendChatMessage(a.Chat, a.Text)

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(result)
}
