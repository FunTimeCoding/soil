package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
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

	project, e := s.resolveProject(a.Project)

	if e != nil {
		return s.captureDetail(e)
	}

	deleted, e := s.service.DeleteRegistryPath(project, a.Path)

	if not_found.Is(e) {
		return response.Fail("no registry repository matches %s", a.Path)
	}

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(deleted)
}
