package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) removeVirtualTag(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(constant.VirtualMachine)

	if f != nil {
		return response.Fail("virtual_machine is required: %v", f)
	}

	tag, g := r.RequireString(constant.Tag)

	if g != nil {
		return response.Fail("tag is required: %v", g)
	}

	result, h := s.client.RemoveVirtualTag(name, tag)

	if h != nil {
		return s.captureFail(h, "tag not removed from virtual machine")
	}

	return response.SuccessAny(convert.VirtualMachine(result))
}
