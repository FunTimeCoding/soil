package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) FindProjects(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.FindProjects,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Projects()

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(result)
}
