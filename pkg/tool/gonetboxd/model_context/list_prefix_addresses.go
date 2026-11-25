package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListPrefixAddresses(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Prefix,
) (*mcp.CallToolResult, error) {
	if a.Prefix == "" {
		return response.Fail("prefix is required")
	}

	result, e := s.service.PrefixAddresses(a.Prefix)

	if e != nil {
		return s.answerMessage(x, e, "addresses not listed")
	}

	return response.SuccessAny(convert.Addresses(result))
}
