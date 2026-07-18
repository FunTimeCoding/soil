package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) setVirtualLabel(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(constant.VirtualMachine)

	if f != nil {
		return response.Fail("virtual_machine is required: %v", f)
	}

	key, g := r.RequireString(constant.Key)

	if g != nil {
		return response.Fail("key is required: %v", g)
	}

	value, h := r.RequireString(constant.Value)

	if h != nil {
		return response.Fail("value is required: %v", h)
	}

	m, resolveFail := s.client.VirtualMachineByName(name)

	if resolveFail != nil {
		return s.captureFail(resolveFail, "virtual machine not found")
	}

	result, setFail := s.store.SetLabel(
		netboxConstant.VirtualMachineAddress,
		m.Identifier,
		key,
		value,
	)

	if setFail != nil {
		return s.captureFail(setFail, "label not set on virtual machine")
	}

	return response.SuccessAny(convert.Label(result))
}
