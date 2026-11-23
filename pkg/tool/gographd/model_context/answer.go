package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) answer(
	x context.Context,
	e error,
) (*mcp.CallToolResult, error) {
	return response.AnswerMessage(x, s.reporter, e, "")
}
