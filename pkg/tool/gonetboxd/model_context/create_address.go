package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createAddress(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	device, f := r.RequireString(constant.Device)

	if f != nil {
		return response.Fail("device is required: %v", f)
	}

	interfaceName, g := r.RequireString(constant.Interface)

	if g != nil {
		return response.Fail("interface is required: %v", g)
	}

	address, h := r.RequireString(constant.Address)

	if h != nil {
		return response.Fail("address is required: %v", h)
	}

	result, e := s.service.CreateAddress(
		device,
		interfaceName,
		address,
		r.GetString(constant.Status, ""),
	)

	if e != nil {
		return s.serviceFail(e, "address not assigned")
	}

	return response.SuccessAny(convert.Address(result))
}
