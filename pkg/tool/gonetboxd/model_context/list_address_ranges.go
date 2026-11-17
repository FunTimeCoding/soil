package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListAddressRanges(
	x context.Context,
	_ mcp.CallToolRequest,
	_ struct{},
) (*mcp.CallToolResult, error) {
	result, e := s.service.InternetAddressRanges()

	if e != nil {
		return s.answerMessage(x, e, "address ranges not retrieved")
	}

	return response.SuccessAny(convert.AddressRanges(result))
}
