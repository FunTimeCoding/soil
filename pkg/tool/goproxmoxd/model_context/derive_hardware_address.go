package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) DeriveHardwareAddress(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DeriveHardwareAddress,
) (*mcp.CallToolResult, error) {
	if a.Identifier == 0 {
		return response.Fail("identifier is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	address, holder, g := s.service.DeriveHardwareAddress(
		instance,
		a.Identifier,
	)

	if g != nil {
		if validation.Is(g) {
			return response.Fail("%s", g)
		}

		return s.captureDetail(g)
	}

	return response.SuccessAny(
		convert.DerivedAddress(instance, a.Identifier, address, holder),
	)
}
