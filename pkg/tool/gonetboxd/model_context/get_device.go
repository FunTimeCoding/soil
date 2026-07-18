package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getDevice(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(parameter.Name)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	d, g := s.client.DeviceByName(name)

	if g != nil {
		return s.captureFail(g, "device not found")
	}

	labels, e := s.store.Labels(
		constant.DeviceAddress,
		d.Identifier,
	)

	if e != nil {
		return s.captureFail(e, "labels not listed for device")
	}

	result := convert.Device(d)
	result.Labels = new(convert.Labels(labels))

	return response.SuccessAny(result)
}
