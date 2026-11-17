package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateInterface(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateInterface,
) (*mcp.CallToolResult, error) {
	if a.Device == "" || a.Name == "" || a.Type == "" {
		return response.Fail("device, name and type are required")
	}

	result, e := s.service.CreateInterface(
		a.Device,
		a.Name,
		a.Type,
		a.PhysicalAddress,
	)

	if e != nil {
		return s.answerMessage(x, e, "interface not created")
	}

	return response.SuccessAny(convert.Interface(result))
}
