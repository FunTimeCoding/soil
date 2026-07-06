package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/netbox-community/go-netbox/v4"
)

func (s *Server) createInterface(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	device, f := r.RequireString(constant.Device)

	if f != nil {
		return response.Fail("device is required: %v", f)
	}

	name, g := r.RequireString(parameter.Name)

	if g != nil {
		return response.Fail("name is required: %v", g)
	}

	interfaceType, h := r.RequireString(constant.Type)

	if h != nil {
		return response.Fail("type is required: %v", h)
	}

	d, i := s.client.DeviceByName(device)

	if i != nil {
		return s.captureFail(i, "device not found")
	}

	result, j := s.client.CreateInterface(
		d,
		name,
		netbox.InterfaceTypeValue(interfaceType),
	)

	if j != nil {
		return s.captureFail(j, "interface not created")
	}

	return response.SuccessAny(convert.Interface(result))
}
