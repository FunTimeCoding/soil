package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureDetail(x context.Context, e error) (*mcp.CallToolResult, error) {
	return response.Answer(x, s.reporter, e)
}
