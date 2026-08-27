package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (s *Server) DeleteRegistryRepository(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.DeleteRegistryRepository,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Path == "" {
		return response.Fail("path is required")
	}

	list, _, e := s.client.ContainerRegistry.ListProjectRegistryRepositories(
		a.Project,
		&gitlab.ListProjectRegistryRepositoriesOptions{
			ListOptions: gitlab.ListOptions{PerPage: 100},
		},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	targets := []string{a.Path, join.Slash([]string{a.Path, "cache"})}
	var deleted []string

	for _, target := range targets {
		for _, repository := range list {
			if repository.Path != target {
				continue
			}

			if _, f := s.client.ContainerRegistry.DeleteRegistryRepository(
				a.Project,
				repository.ID,
			); f != nil {
				return s.captureDetail(f)
			}

			deleted = append(deleted, target)

			break
		}
	}

	if len(deleted) == 0 {
		return response.Fail("no registry repository matches %s", a.Path)
	}

	return response.SuccessAny(map[string]any{"deleted": deleted})
}
