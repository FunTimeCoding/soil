package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createDeviceType(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	model, f := r.RequireString(constant.Model)

	if f != nil {
		return response.Fail("model is required: %v", f)
	}

	manufacturer, g := r.RequireString(constant.Manufacturer)

	if g != nil {
		return response.Fail("manufacturer is required: %v", g)
	}

	m, h := s.client.ManufacturerByName(manufacturer)

	if h != nil {
		return s.captureFail(h, "manufacturer not found")
	}

	result, i := s.client.CreateDeviceType(model, m)

	if i != nil {
		return s.captureFail(i, "device type not created")
	}

	return response.SuccessAny(convert.DeviceType(result))
}
