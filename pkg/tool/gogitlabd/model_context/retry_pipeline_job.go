package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) RetryPipelineJob(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.RetryPipelineJob,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Job == 0 {
		return response.Fail("job is required")
	}

	v, _, e := s.client.Jobs.RetryJob(a.Project, a.Job)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
