package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_assignment"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createConfigurationContext(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, e := r.RequireString(generativeConstant.ParameterName)

	if e != nil {
		return response.Fail("name is required: %v", e)
	}

	payload, okay := r.GetArguments()[constant.Payload].(map[string]any)

	if !okay {
		return response.Fail("payload is required and must be an object")
	}

	a := assignmentArgument(r)

	if a == nil {
		a = configuration_assignment.New()
	}

	result, f := s.service.CreateConfigurationContext(
		name,
		int32(
			r.GetInt(constant.Weight, int(constant.ConfigurationContextWeight)),
		),
		payload,
		a,
	)

	if not_found.Is(f) {
		return response.Fail("%v", f)
	}

	if f != nil {
		return s.captureFail(f, "configuration context not created")
	}

	return response.SuccessAny(convert.ConfigurationContext(result))
}
