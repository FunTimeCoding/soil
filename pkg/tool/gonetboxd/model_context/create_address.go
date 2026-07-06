package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createAddress(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	device, f := r.RequireString(constant.Device)

	if f != nil {
		return response.Fail("device is required: %v", f)
	}

	interfaceName, g := r.RequireString(constant.Interface)

	if g != nil {
		return response.Fail("interface is required: %v", g)
	}

	address, h := r.RequireString(constant.Address)

	if h != nil {
		return response.Fail("address is required: %v", h)
	}

	d, i := s.client.DeviceByName(device)

	if i != nil {
		return s.captureFail(i, "device not found")
	}

	iface, j := s.client.DeviceInterfaceByName(d, interfaceName)

	if j != nil {
		return s.captureFail(j, "interface not found on device")
	}

	result, k := s.client.CreateAddress(iface.Identifier, address)

	if k != nil {
		return s.captureFail(k, "address not assigned")
	}

	return response.SuccessAny(convert.Address(result))
}
