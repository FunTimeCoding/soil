package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateAddressRange(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateAddressRange,
) (*mcp.CallToolResult, error) {
	if a.Start == "" || a.End == "" {
		return response.Fail("start and end are required")
	}

	result, e := s.service.CreateInternetAddressRange(
		a.Start,
		a.End,
		a.Status,
		a.Description,
	)

	if e != nil {
		return s.answerMessage(x, e, "address range not created")
	}

	return response.SuccessAny(convert.AddressRange(result))
}
