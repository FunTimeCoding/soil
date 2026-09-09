package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListPipelineJobs(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListPipelineJobs,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Pipeline == 0 {
		return response.Fail("pipeline is required")
	}

	project, e := s.resolveProject(a.Project)

	if e != nil {
		return s.captureDetail(e)
	}

	v, e := s.client.PipelineJobs(project, a.Pipeline)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
