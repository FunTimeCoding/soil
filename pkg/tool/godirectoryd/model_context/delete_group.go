package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteGroup(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Group,
) (*mcp.CallToolResult, error) {
	if e := s.service.DeleteGroup(a.Name); e != nil {
		return s.captureDetail(e)
	}

	return response.Success(fmt.Sprintf("deleted group: %s", a.Name))
}
