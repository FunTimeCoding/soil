package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListProjects(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListProjects,
) (*mcp.CallToolResult, error) {
	v, e := s.service.Projects(a.Search)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
