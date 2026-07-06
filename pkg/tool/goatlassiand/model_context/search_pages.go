package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) searchPages(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	query, f := r.RequireString(parameter.Query)

	if f != nil {
		return response.Fail("query is required: %v", f)
	}

	result, g := s.confluence.Search(query)

	if g != nil {
		return s.captureDetail(g)
	}

	return response.SuccessAny(result)
}
