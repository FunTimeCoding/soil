package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateBranch(
	x context.Context,
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

	instance, project, e := s.resolveProject(x, a.Project)

	if e != nil {
		return s.captureDetail(e)
	}

	v, e := s.service.CreateBranch(instance, project, a.Branch, a.Reference)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
