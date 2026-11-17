package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/types/machine_change"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) UpdateVirtualMachine(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.UpdateVirtualMachine,
) (*mcp.CallToolResult, error) {
	if a.VirtualMachine == "" {
		return response.Fail("virtual machine is required")
	}

	c := machine_change.New()
	c.Name = filled(a.Name)
	c.PrimaryAddress = filled(a.PrimaryAddress)
	c.Platform = filled(a.Platform)
	c.Role = filled(a.Role)
	c.Tenant = filled(a.Tenant)
	c.Status = filled(a.Status)
	c.CustomFields = a.CustomFields
	c.LocalContext = a.LocalContext

	if a.Cores > 0 {
		c.Cores = new(a.Cores)
	}

	if a.Memory > 0 {
		c.Memory = new(int32(a.Memory))
	}

	result, e := s.service.UpdateVirtualMachine(a.VirtualMachine, c)

	if e != nil {
		return s.answerMessage(x, e, "virtual machine not updated")
	}

	return response.SuccessAny(convert.VirtualMachine(result))
}
