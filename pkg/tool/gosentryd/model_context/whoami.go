package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Whoami(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.Whoami,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Whoami()

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(result)
}
