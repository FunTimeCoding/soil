package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateDevice(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateDevice,
) (*mcp.CallToolResult, error) {
	if a.Name == "" || a.Role == "" || a.Type == "" || a.Site == "" {
		return response.Fail("name, role, type and site are required")
	}

	result, e := s.service.CreateDevice(
		a.Name,
		a.Role,
		a.Type,
		a.Site,
		a.Tenant,
		a.Tags,
	)

	if e != nil {
		return s.answerMessage(x, e, "device not created")
	}

	return response.SuccessAny(convert.Device(result))
}
