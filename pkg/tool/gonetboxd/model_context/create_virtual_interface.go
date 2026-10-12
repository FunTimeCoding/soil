package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createVirtualInterface(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	vmName, f := r.RequireString(constant.VirtualMachine)

	if f != nil {
		return response.Fail("virtual_machine is required: %v", f)
	}

	name, g := r.RequireString(generativeConstant.ParameterName)

	if g != nil {
		return response.Fail("name is required: %v", g)
	}

	result, e := s.service.CreateVirtualInterface(vmName, name)

	if e != nil {
		return s.serviceFail(e, "virtual interface not created")
	}

	return response.SuccessAny(convert.VirtualInterface(result))
}
