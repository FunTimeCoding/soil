package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) FindReleases(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.FindReleases,
) (*mcp.CallToolResult, error) {
	result, e := s.client.Releases(
		s.organization,
		a.Query,
		a.Limit,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
