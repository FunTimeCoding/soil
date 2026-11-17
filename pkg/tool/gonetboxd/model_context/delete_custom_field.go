package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) DeleteCustomField(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Name,
) (*mcp.CallToolResult, error) {
	if a.Name == "" {
		return response.Fail("name is required")
	}

	e := s.service.DeleteCustomField(a.Name)

	if e != nil {
		return s.answerMessage(x, e, "custom field not deleted")
	}

	return response.Success("custom field deleted")
}
