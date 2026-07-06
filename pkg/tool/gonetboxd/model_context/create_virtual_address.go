package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createVirtualAddress(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	vmName, f := r.RequireString(constant.VirtualMachine)

	if f != nil {
		return response.Fail("virtual_machine is required: %v", f)
	}

	interfaceName, g := r.RequireString(constant.Interface)

	if g != nil {
		return response.Fail("interface is required: %v", g)
	}

	address, h := r.RequireString(constant.Address)

	if h != nil {
		return response.Fail("address is required: %v", h)
	}

	vm, i := s.client.VirtualMachineByName(vmName)

	if i != nil {
		return s.captureFail(i, "virtual machine not found")
	}

	iface, j := s.client.VirtualMachineInterfaceByName(vm, interfaceName)

	if j != nil {
		return s.captureFail(j, "interface not found on virtual machine")
	}

	result, k := s.client.CreateVirtualAddress(iface.GetId(), address)

	if k != nil {
		return s.captureFail(k, "address not assigned")
	}

	return response.SuccessAny(convert.Address(result))
}
