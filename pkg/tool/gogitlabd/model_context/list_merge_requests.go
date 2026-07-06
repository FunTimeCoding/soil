package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (s *Server) ListMergeRequests(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListMergeRequests,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	o := &gitlab.ListProjectMergeRequestsOptions{
		ListOptions: gitlab.ListOptions{PerPage: 20},
	}

	if a.State != "" {
		o.State = &a.State
	}

	v, _, e := s.client.MergeRequests.ListProjectMergeRequests(
		a.Project,
		o,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
