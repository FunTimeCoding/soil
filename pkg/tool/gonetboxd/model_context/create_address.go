package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateAddress(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateAddress,
) (*mcp.CallToolResult, error) {
	if a.Device == "" || a.Interface == "" || a.Address == "" {
		return response.Fail("device, interface and address are required")
	}

	result, e := s.service.CreateAddress(
		a.Device,
		a.Interface,
		a.Address,
		a.Status,
	)

	if e != nil {
		return s.answerMessage(x, e, "address not assigned")
	}

	return response.SuccessAny(convert.Address(result))
}
