package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Whoami(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.Whoami,
) (*mcp.CallToolResult, error) {
	result, e := s.client.Whoami()

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
