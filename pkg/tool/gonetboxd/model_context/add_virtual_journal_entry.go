package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) AddVirtualJournalEntry(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.VirtualJournalEntry,
) (*mcp.CallToolResult, error) {
	if a.VirtualMachine == "" || a.Comments == "" {
		return response.Fail("virtual machine and comments are required")
	}

	result, e := s.service.AddVirtualJournalEntry(
		a.VirtualMachine,
		a.Kind,
		a.Comments,
	)

	if e != nil {
		return s.answerMessage(
			x,
			e,
			"journal entry not added to virtual machine",
		)
	}

	return response.SuccessAny(result)
}
