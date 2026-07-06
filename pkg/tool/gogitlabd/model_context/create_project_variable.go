package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (s *Server) CreateProjectVariable(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateProjectVariable,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Key == "" {
		return response.Fail("key is required")
	}

	if a.Value == "" {
		return response.Fail("value is required")
	}

	v, _, e := s.client.ProjectVariables.CreateVariable(
		a.Project,
		&gitlab.CreateProjectVariableOptions{
			Key:       &a.Key,
			Value:     &a.Value,
			Protected: &a.Protected,
			Masked:    &a.Masked,
		},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
