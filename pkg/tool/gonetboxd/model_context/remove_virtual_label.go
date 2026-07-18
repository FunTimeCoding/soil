package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) removeVirtualLabel(
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

	m, resolveFail := s.client.VirtualMachineByName(name)

	if resolveFail != nil {
		return s.captureFail(resolveFail, "virtual machine not found")
	}

	if e := s.store.RemoveLabel(
		netboxConstant.VirtualMachineAddress,
		m.Identifier,
		key,
	); e != nil {
		return s.captureFail(e, "label not removed from virtual machine")
	}

	return response.Success("label removed")
}
