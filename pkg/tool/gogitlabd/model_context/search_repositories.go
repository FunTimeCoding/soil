package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SearchRepositories(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SearchRepositories,
) (*mcp.CallToolResult, error) {
	if a.Query == "" {
		return response.Fail("query is required")
	}

	v, e := s.client.SearchProject(a.Query)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
