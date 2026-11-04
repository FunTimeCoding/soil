package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) revocationList(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Authority,
) (*mcp.CallToolResult, error) {
	result, e := s.service.RevocationList(a.Authority)

	if e != nil {
		return s.answer(x, e)
	}

	return response.Success("%s", string(result))
}
