package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (s *Server) SearchRepositories(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SearchRepositories,
) (*mcp.CallToolResult, error) {
	if a.Query == "" {
		return response.Fail("query is required")
	}

	v, _, e := s.client.Search.Projects(
		a.Query,
		&gitlab.SearchOptions{ListOptions: gitlab.ListOptions{PerPage: 20}},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
