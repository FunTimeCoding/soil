package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListPipelines(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListPipelines,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	project, e := s.resolveProject(a.Project)

	if e != nil {
		return s.captureDetail(e)
	}

	v, e := s.client.Pipelines(project, a.Reference, a.Status, 20)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
