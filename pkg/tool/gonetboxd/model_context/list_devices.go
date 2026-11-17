package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListDevices(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListDevices,
) (*mcp.CallToolResult, error) {
	result, grouped, e := s.service.Devices(a.Query)

	if e != nil {
		return s.answerMessage(x, e, "devices not listed")
	}

	return response.SuccessAny(convert.LabeledDevices(result, grouped))
}
