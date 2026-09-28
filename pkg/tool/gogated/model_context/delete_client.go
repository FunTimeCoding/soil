package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteClient(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Identifier,
) (*mcp.CallToolResult, error) {
	if _, e := s.service.Client(a.Identifier); e != nil {
		return s.captureDetail(e)
	}

	if e := s.service.DeleteClient(a.Identifier); e != nil {
		return s.captureDetail(e)
	}

	return response.Success(fmt.Sprintf("deleted client: %s", a.Identifier))
}
