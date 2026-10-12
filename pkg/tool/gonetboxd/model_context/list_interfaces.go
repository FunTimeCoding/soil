package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listInterfaces(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(constant.ParameterName)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	result, e := s.service.DeviceInterfaces(name)

	if e != nil {
		return s.serviceFail(e, "interfaces not retrieved")
	}

	return response.SuccessAny(convert.Interfaces(result))
}
