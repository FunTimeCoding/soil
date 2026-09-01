package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) revealPassword(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Identifier,
) (*mcp.CallToolResult, error) {
	password, found := s.service.Reveal(a.Identifier)

	if !found {
		return response.Fail("entry not found: %s", a.Identifier)
	}

	return response.Success(password)
}
