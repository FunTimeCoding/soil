package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (s *Server) UpdateProjectVariable(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.UpdateProjectVariable,
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

	v, _, e := s.client.ProjectVariables.UpdateVariable(
		a.Project,
		a.Key,
		&gitlab.UpdateProjectVariableOptions{
			Value:     &a.Value,
			Protected: &a.Protected,
			Masked:    &a.Masked,
			Raw:       new(!a.Expand),
		},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
