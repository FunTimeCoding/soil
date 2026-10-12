package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createInterface(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	device, f := r.RequireString(constant.Device)

	if f != nil {
		return response.Fail("device is required: %v", f)
	}

	name, g := r.RequireString(generativeConstant.ParameterName)

	if g != nil {
		return response.Fail("name is required: %v", g)
	}

	interfaceType, h := r.RequireString(constant.Type)

	if h != nil {
		return response.Fail("type is required: %v", h)
	}

	result, e := s.service.CreateInterface(
		device,
		name,
		interfaceType,
		r.GetString(constant.PhysicalAddress, ""),
	)

	if e != nil {
		return s.serviceFail(e, "interface not created")
	}

	return response.SuccessAny(convert.Interface(result))
}
