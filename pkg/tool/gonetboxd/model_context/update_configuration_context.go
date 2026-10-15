package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/types/context_change"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) updateConfigurationContext(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, e := r.RequireString(generativeConstant.ParameterName)

	if e != nil {
		return response.Fail("name is required: %v", e)
	}

	c := context_change.New()
	arguments := r.GetArguments()

	if payload, okay := arguments[constant.Payload].(map[string]any); okay {
		c.Payload = payload
	}

	if _, okay := arguments[constant.Weight]; okay {
		c.Weight = new(int32(r.GetInt(constant.Weight, 0)))
	}

	c.Assignment = assignmentArgument(r)
	result, f := s.service.UpdateConfigurationContext(name, c)

	if not_found.Is(f) {
		return response.Fail("%v", f)
	}

	if f != nil {
		return s.captureFail(f, "configuration context not updated")
	}

	return response.SuccessAny(convert.ConfigurationContext(result))
}
