package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listDevices(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	query := r.GetString(parameter.Query, "")
	var result []*device.Device
	var e error

	if query != "" {
		result, e = s.client.DevicesByMatch(query)
	} else {
		result, e = s.client.Devices()
	}

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(convert.Devices(result))
}
