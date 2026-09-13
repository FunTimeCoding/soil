package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) setPassword(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SetPassword,
) (*mcp.CallToolResult, error) {
	if e := s.service.SetPassword(a.Account, a.Password); e != nil {
		return s.captureDetail(e)
	}

	return response.Success(fmt.Sprintf("password set: %s", a.Account))
}
