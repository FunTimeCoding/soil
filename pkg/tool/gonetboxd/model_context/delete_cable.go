package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) DeleteCable(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Identifier,
) (*mcp.CallToolResult, error) {
	if a.Identifier <= 0 {
		return response.Fail("identifier is required")
	}

	if e := s.service.DeleteCable(int32(a.Identifier)); e != nil {
		return s.answerMessage(x, e, "cable not deleted")
	}

	return response.Success("cable deleted")
}
