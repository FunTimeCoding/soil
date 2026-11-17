package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreatePrefix(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreatePrefix,
) (*mcp.CallToolResult, error) {
	if a.Prefix == "" {
		return response.Fail("prefix is required")
	}

	p, e := s.service.CreatePrefix(a.Prefix, a.Site, a.Description)

	if e != nil {
		return s.answerMessage(x, e, "prefix not created")
	}

	return response.SuccessAny(convert.Prefix(p))
}
