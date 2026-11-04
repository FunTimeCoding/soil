package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) answerWith(
	x context.Context,
	e error,
	message string,
) (*mcp.CallToolResult, error) {
	return response.AnswerWith(x, s.reporter, e, message)
}
