package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listDevices(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, grouped, e := s.service.Devices(
		r.GetString(constant.ParameterQuery, ""),
	)

	if e != nil {
		return s.serviceFail(e, "devices not listed")
	}

	return response.SuccessAny(convert.LabeledDevices(result, grouped))
}
