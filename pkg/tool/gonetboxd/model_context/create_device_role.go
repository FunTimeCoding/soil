package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createDeviceRole(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(generativeConstant.ParameterName)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	result, g := s.client.CreateDeviceRole(
		name,
		r.GetBool(constant.VirtualMachineRole, false),
	)

	if g != nil {
		return s.captureFail(g, "device role not created")
	}

	return response.SuccessAny(convert.DeviceRole(result))
}
