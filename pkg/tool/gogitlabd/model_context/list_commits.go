package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListCommits(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListCommits,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	instance, project, e := s.resolveProject(x, a.Project)

	if e != nil {
		return s.captureDetail(e)
	}

	v, e := s.service.ListCommits(instance, project, a.Reference, 0)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
