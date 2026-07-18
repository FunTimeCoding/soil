package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listVirtualJournalEntries(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(constant.VirtualMachine)

	if f != nil {
		return response.Fail("virtual_machine is required: %v", f)
	}

	result, g := s.client.VirtualJournalEntries(
		name,
		int32(r.GetInt(parameter.Limit, 0)),
		int32(r.GetInt(parameter.Offset, 0)),
	)

	if g != nil {
		return s.captureFail(
			g,
			"journal entries not listed for virtual machine",
		)
	}

	return response.SuccessAny(result)
}
