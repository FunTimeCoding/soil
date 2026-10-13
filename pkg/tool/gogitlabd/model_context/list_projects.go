package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListProjects(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListProjects,
) (*mcp.CallToolResult, error) {
	instance, e := s.resolveInstance(x)

	if e != nil {
		return s.captureDetail(e)
	}

	v, e := s.service.Projects(instance, a.Search)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
