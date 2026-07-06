package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetPipeline(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetPipeline,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Pipeline == 0 {
		return response.Fail("pipeline is required")
	}

	v, _, e := s.client.Pipelines.GetPipeline(a.Project, a.Pipeline)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
