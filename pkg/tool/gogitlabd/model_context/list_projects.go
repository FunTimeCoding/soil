package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (s *Server) ListProjects(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListProjects,
) (*mcp.CallToolResult, error) {
	o := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{PerPage: 20},
	}

	if a.Search != "" {
		o.Search = &a.Search
	}

	v, _, e := s.client.Projects.ListProjects(o)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
