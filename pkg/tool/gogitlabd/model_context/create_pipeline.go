package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (s *Server) CreatePipeline(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CreatePipeline,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Reference == "" {
		return response.Fail("reference is required")
	}

	project, e := s.resolveProject(a.Project)

	if e != nil {
		return s.captureDetail(e)
	}

	options := &gitlab.CreatePipelineOptions{Ref: &a.Reference}

	if len(a.Variables) > 0 {
		vars := make([]*gitlab.PipelineVariableOptions, 0, len(a.Variables))

		for key, value := range a.Variables {
			k := key
			l := value
			vars = append(
				vars,
				&gitlab.PipelineVariableOptions{
					Key:   &k,
					Value: &l,
				},
			)
		}

		options.Variables = &vars
	}

	v, _, e := s.client.Pipelines.CreatePipeline(
		project,
		options,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
