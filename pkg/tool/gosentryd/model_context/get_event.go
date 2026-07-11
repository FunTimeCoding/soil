package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetEvent(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetEvent,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Identifier == "" {
		return response.Fail("identifier is required")
	}

	result, e := s.client.Event(
		s.organization,
		a.Project,
		a.Identifier,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
