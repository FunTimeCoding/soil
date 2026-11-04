package model_context

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) serviceFail(
	x context.Context,
	e error,
	message string,
) (*mcp.CallToolResult, error) {
	return s.answerWith(x, e, message)
}
