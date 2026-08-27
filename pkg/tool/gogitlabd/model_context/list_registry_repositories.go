package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (s *Server) ListRegistryRepositories(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListRegistryRepositories,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	v, _, e := s.client.ContainerRegistry.ListProjectRegistryRepositories(
		a.Project,
		&gitlab.ListProjectRegistryRepositoriesOptions{
			ListOptions: gitlab.ListOptions{PerPage: 100},
		},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
