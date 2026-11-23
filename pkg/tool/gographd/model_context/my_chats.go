package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gographd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) MyChats(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.MyChats,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Chats(a.Limit)

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(result)
}
