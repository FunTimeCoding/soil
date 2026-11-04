package model_context

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) answer(x context.Context, e error) (*mcp.CallToolResult, error) {
	return s.answerWith(x, e, "")
}
