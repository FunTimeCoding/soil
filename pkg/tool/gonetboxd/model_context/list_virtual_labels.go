package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listVirtualLabels(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(constant.VirtualMachine)

	if f != nil {
		return response.Fail("virtual_machine is required: %v", f)
	}

	m, resolveFail := s.client.VirtualMachineByName(name)

	if resolveFail != nil {
		return s.captureFail(resolveFail, "virtual machine not found")
	}

	labels, e := s.store.Labels(
		netboxConstant.VirtualMachineAddress,
		m.Identifier,
	)

	if e != nil {
		return s.captureFail(e, "labels not listed for virtual machine")
	}

	return response.SuccessAny(convert.Labels(labels))
}
