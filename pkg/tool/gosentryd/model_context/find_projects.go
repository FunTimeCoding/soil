package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) FindProjects(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.FindProjects,
) (*mcp.CallToolResult, error) {
	result, e := s.client.OrganizationProjects(s.organization)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
