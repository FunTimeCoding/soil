package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_assignment"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateConfigurationContext(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ConfigurationContext,
) (*mcp.CallToolResult, error) {
	if a.Name == "" {
		return response.Fail("name is required")
	}

	if a.Payload == nil {
		return response.Fail("payload is required and must be an object")
	}

	assignment := assignmentArgument(a)

	if assignment == nil {
		assignment = configuration_assignment.New()
	}

	weight := constant.ConfigurationContextWeight

	if a.Weight != nil {
		weight = int32(*a.Weight)
	}

	result, e := s.service.CreateConfigurationContext(
		a.Name,
		weight,
		*a.Payload,
		assignment,
	)

	if e != nil {
		return s.answerMessage(x, e, "configuration context not created")
	}

	return response.SuccessAny(convert.ConfigurationContext(result))
}
