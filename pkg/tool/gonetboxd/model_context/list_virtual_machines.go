package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listVirtualMachines(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, grouped, e := s.service.VirtualMachines()

	if e != nil {
		return s.serviceFail(e, "virtual machines not listed")
	}

	return response.SuccessAny(convert.LabeledVirtualMachines(result, grouped))
}
