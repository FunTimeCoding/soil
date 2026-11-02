package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) answerWith(
	e error,
	message string,
) (*mcp.CallToolResult, error) {
	return response.AnswerWith(s.reporter, e, message)
}
