package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListPipelines(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListPipelines,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	instance, project, e := s.resolveProject(x, a.Project)

	if e != nil {
		return s.captureDetail(e)
	}

	v, e := s.service.Pipelines(instance, project, a.Reference, a.Status, 20)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
