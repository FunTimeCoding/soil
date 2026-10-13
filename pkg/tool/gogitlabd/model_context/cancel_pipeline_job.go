package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	gitLabResponse "github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CancelPipelineJob(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CancelPipelineJob,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Job == 0 {
		return response.Fail("job is required")
	}

	instance, project, e := s.resolveProject(x, a.Project)

	if e != nil {
		return s.captureDetail(e)
	}

	v, e := s.service.CancelJob(instance, project, a.Job)

	if e != nil {
		return s.captureDetail(e)
	}

	if a.Raw {
		return response.SuccessAnyRaw(v)
	}

	return response.SuccessAny(gitLabResponse.NewJob(v))
}
