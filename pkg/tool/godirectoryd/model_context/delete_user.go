package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteUser(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Account,
) (*mcp.CallToolResult, error) {
	if e := s.service.DeleteUser(a.Account); e != nil {
		return s.captureDetail(e)
	}

	return response.Success(fmt.Sprintf("deleted user: %s", a.Account))
}
