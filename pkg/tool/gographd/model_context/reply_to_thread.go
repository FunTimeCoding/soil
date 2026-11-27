package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gographd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ReplyToThread(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ReplyToThread,
) (*mcp.CallToolResult, error) {
	result, e := s.service.ReplyToThread(
		a.Team,
		a.Channel,
		a.Message,
		a.Text,
		a.Mention,
	)

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(result)
}
