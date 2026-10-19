package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateTab(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.CreateTab,
) (*mcp.CallToolResult, error) {
	v, e := s.client.CreateTab()

	if e != nil {
		return s.answer(e)
	}

	return response.SuccessAny(v)
}
