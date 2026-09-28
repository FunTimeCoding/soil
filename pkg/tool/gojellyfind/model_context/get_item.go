package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetItem(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetItem,
) (*mcp.CallToolResult, error) {
	if a.Identifier == "" {
		return response.Fail("id is required")
	}

	result, e := s.client.Item(a.Identifier)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
