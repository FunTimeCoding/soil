package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listCustomFields(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, e := s.service.CustomFields()

	if e != nil {
		return s.captureFail(x, e, "custom fields not listed")
	}

	return response.SuccessAny(convert.CustomFields(result))
}
