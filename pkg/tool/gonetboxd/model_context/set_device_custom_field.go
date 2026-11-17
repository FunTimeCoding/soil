package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SetDeviceCustomField(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DeviceCustomField,
) (*mcp.CallToolResult, error) {
	if a.Device == "" || a.Field == "" {
		return response.Fail("device and field are required")
	}

	values := a.Values

	if values == nil {
		values = []string{}
	}

	result, e := s.service.SetDeviceCustomFields(
		a.Device,
		map[string]any{a.Field: values},
	)

	if e != nil {
		return s.answerMessage(x, e, "custom field not set")
	}

	return response.SuccessAny(convert.Device(result))
}
