package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) bumpDecision(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.BumpDecision,
) (*mcp.CallToolResult, error) {
	s.service.Bump(uint(a.Identifier))

	return response.Success("marked as needed")
}
