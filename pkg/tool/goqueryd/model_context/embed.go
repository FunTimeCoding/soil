package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) embed(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Embed(x)

	if unavailable.Is(e) {
		return response.Fail("%v", e)
	}

	if e != nil {
		return s.captureFail(e, "embed failed")
	}

	return response.Success(
		"embedded %d documents (%d chunks)",
		result.Documents,
		result.Chunks,
	)
}
