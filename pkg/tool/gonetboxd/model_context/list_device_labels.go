package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listDeviceLabels(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	device, f := r.RequireString(constant.Device)

	if f != nil {
		return response.Fail("device is required: %v", f)
	}

	d, resolveFail := s.client.DeviceByName(device)

	if resolveFail != nil {
		return s.captureFail(resolveFail, "device not found")
	}

	labels, e := s.store.Labels(
		netboxConstant.DeviceAddress,
		d.Identifier,
	)

	if e != nil {
		return s.captureFail(e, "labels not listed for device")
	}

	return response.SuccessAny(convert.Labels(labels))
}
