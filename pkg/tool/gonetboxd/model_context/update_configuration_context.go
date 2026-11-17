package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/types/context_change"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) UpdateConfigurationContext(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ConfigurationContext,
) (*mcp.CallToolResult, error) {
	if a.Name == "" {
		return response.Fail("name is required")
	}

	c := context_change.New()

	if a.Payload != nil {
		c.Payload = *a.Payload
	}

	if a.Weight != nil {
		c.Weight = new(int32(*a.Weight))
	}

	c.Assignment = assignmentArgument(a)
	result, e := s.service.UpdateConfigurationContext(a.Name, c)

	if e != nil {
		return s.answerMessage(x, e, "configuration context not updated")
	}

	return response.SuccessAny(convert.ConfigurationContext(result))
}
