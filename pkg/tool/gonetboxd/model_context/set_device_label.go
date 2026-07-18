package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) setDeviceLabel(
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

	value, h := r.RequireString(constant.Value)

	if h != nil {
		return response.Fail("value is required: %v", h)
	}

	d, resolveFail := s.client.DeviceByName(device)

	if resolveFail != nil {
		return s.captureFail(resolveFail, "device not found")
	}

	result, setFail := s.store.SetLabel(
		netboxConstant.DeviceAddress,
		d.Identifier,
		key,
		value,
	)

	if setFail != nil {
		return s.captureFail(setFail, "label not set on device")
	}

	return response.SuccessAny(convert.Label(result))
}
