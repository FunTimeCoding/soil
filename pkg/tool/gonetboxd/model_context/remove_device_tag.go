package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) removeDeviceTag(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	device, f := r.RequireString(constant.Device)

	if f != nil {
		return response.Fail("device is required: %v", f)
	}

	tag, g := r.RequireString(constant.Tag)

	if g != nil {
		return response.Fail("tag is required: %v", g)
	}

	result, h := s.client.RemoveTag(device, tag)

	if h != nil {
		return s.captureFail(h, "tag not removed from device")
	}

	return response.SuccessAny(convert.Device(result))
}
