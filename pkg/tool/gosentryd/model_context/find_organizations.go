package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) FindOrganizations(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.FindOrganizations,
) (*mcp.CallToolResult, error) {
	result, e := s.client.Organizations()

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
