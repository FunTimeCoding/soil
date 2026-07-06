package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (s *Server) CreateBranch(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateBranch,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Branch == "" {
		return response.Fail("branch is required")
	}

	if a.Reference == "" {
		return response.Fail("reference is required")
	}

	v, _, e := s.client.Branches.CreateBranch(
		a.Project,
		&gitlab.CreateBranchOptions{
			Branch: &a.Branch,
			Ref:    &a.Reference,
		},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
