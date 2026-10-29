package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listAuthorities(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ struct{},
) (*mcp.CallToolResult, error) {
	result, e := s.service.Authorities()

	if e != nil {
		return s.answer(e)
	}

	return response.SuccessAny(convert.Authorities(result))
}
