package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetConfigurationContext(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Name,
) (*mcp.CallToolResult, error) {
	if a.Name == "" {
		return response.Fail("name is required")
	}

	result, e := s.service.ConfigurationContextByName(a.Name)

	if e != nil {
		return s.answerMessage(x, e, "configuration context not read")
	}

	return response.SuccessAny(convert.ConfigurationContext(result))
}
