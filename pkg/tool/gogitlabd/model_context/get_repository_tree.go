package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (s *Server) GetRepositoryTree(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetRepositoryTree,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	o := &gitlab.ListTreeOptions{
		ListOptions: gitlab.ListOptions{PerPage: 100},
	}

	if a.Path != "" {
		o.Path = &a.Path
	}

	if a.Reference != "" {
		o.Ref = &a.Reference
	}

	if a.Recursive {
		o.Recursive = &a.Recursive
	}

	v, _, e := s.client.Repositories.ListTree(a.Project, o)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
