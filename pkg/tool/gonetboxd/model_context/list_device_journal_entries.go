package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListDeviceJournalEntries(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DeviceJournalPage,
) (*mcp.CallToolResult, error) {
	if a.Device == "" {
		return response.Fail("device is required")
	}

	result, e := s.service.DeviceJournalEntries(
		a.Device,
		int32(a.Limit),
		int32(a.Offset),
	)

	if e != nil {
		return s.answerMessage(x, e, "journal entries not listed for device")
	}

	return response.SuccessAny(result)
}
