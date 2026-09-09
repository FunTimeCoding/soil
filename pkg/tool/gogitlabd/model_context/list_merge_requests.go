package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListMergeRequests(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListMergeRequests,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	project, e := s.resolveProject(a.Project)

	if e != nil {
		return s.captureDetail(e)
	}

	v, e := s.client.ProjectMergeRequests(project, a.State)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
