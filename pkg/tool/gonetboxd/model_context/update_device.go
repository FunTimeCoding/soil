package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/types/device_change"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) UpdateDevice(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.UpdateDevice,
) (*mcp.CallToolResult, error) {
	if a.Device == "" {
		return response.Fail("device is required")
	}

	c := device_change.New()
	c.Name = filled(a.Name)
	c.PrimaryAddress = filled(a.PrimaryAddress)
	c.Location = filled(a.Location)
	c.Platform = filled(a.Platform)
	c.Type = filled(a.Type)
	c.Role = filled(a.Role)
	c.Site = filled(a.Site)
	c.Tenant = filled(a.Tenant)
	c.Status = filled(a.Status)
	c.Serial = filled(a.Serial)
	c.Description = filled(a.Description)
	c.CustomFields = a.CustomFields
	c.LocalContext = a.LocalContext
	result, e := s.service.UpdateDevice(a.Device, c)

	if e != nil {
		return s.answerMessage(x, e, "device not updated")
	}

	return response.SuccessAny(convert.Device(result))
}
