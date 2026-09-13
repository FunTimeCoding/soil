package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addMember(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Member,
) (*mcp.CallToolResult, error) {
	result, e := s.service.AddMember(a.Name, a.Account)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
