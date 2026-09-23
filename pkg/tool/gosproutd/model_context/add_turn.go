package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addTurn(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.AddTurn,
) (*mcp.CallToolResult, error) {
	_, e := s.service.AddTurn(
		uint(a.Identifier),
		constant.AuthorSession,
		a.Content,
	)

	if e != nil {
		return response.Fail("%s", e)
	}

	return response.Success("turn added")
}
