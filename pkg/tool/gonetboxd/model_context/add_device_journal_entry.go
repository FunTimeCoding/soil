package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) AddDeviceJournalEntry(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DeviceJournalEntry,
) (*mcp.CallToolResult, error) {
	if a.Device == "" || a.Comments == "" {
		return response.Fail("device and comments are required")
	}

	result, e := s.service.AddDeviceJournalEntry(a.Device, a.Kind, a.Comments)

	if e != nil {
		return s.answerMessage(x, e, "journal entry not added to device")
	}

	return response.SuccessAny(result)
}
