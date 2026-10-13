package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	gitLabResponse "github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListPipelineJobs(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListPipelineJobs,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Pipeline == 0 {
		return response.Fail("pipeline is required")
	}

	instance, project, e := s.resolveProject(x, a.Project)

	if e != nil {
		return s.captureDetail(e)
	}

	v, e := s.service.PipelineJobs(instance, project, a.Pipeline)

	if e != nil {
		return s.captureDetail(e)
	}

	if a.Raw {
		return response.SuccessAnyRaw(v)
	}

	return response.SuccessAny(gitLabResponse.NewJobSlice(v))
}
