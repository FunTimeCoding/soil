package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) FindReleases(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.FindReleases,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Releases(a.Query, a.Limit)

	if e != nil {
		return s.answer(e)
	}

	return s.succeed(result)
}
