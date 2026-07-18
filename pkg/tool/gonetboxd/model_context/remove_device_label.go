package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) removeDeviceLabel(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	device, f := r.RequireString(constant.Device)

	if f != nil {
		return response.Fail("device is required: %v", f)
	}

	key, g := r.RequireString(constant.Key)

	if g != nil {
		return response.Fail("key is required: %v", g)
	}

	d, resolveFail := s.client.DeviceByName(device)

	if resolveFail != nil {
		return s.captureFail(resolveFail, "device not found")
	}

	if e := s.store.RemoveLabel(
		netboxConstant.DeviceAddress,
		d.Identifier,
		key,
	); e != nil {
		return s.captureFail(e, "label not removed from device")
	}

	return response.Success("label removed")
}
