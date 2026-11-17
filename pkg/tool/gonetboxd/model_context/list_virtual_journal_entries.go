package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListVirtualJournalEntries(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.VirtualJournalPage,
) (*mcp.CallToolResult, error) {
	if a.VirtualMachine == "" {
		return response.Fail("virtual machine is required")
	}

	result, e := s.service.VirtualJournalEntries(
		a.VirtualMachine,
		int32(a.Limit),
		int32(a.Offset),
	)

	if e != nil {
		return s.answerMessage(
			x,
			e,
			"journal entries not listed for virtual machine",
		)
	}

	return response.SuccessAny(result)
}
