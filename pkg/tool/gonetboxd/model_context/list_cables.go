package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListCables(
	x context.Context,
	_ mcp.CallToolRequest,
	_ struct{},
) (*mcp.CallToolResult, error) {
	result, e := s.service.Cables()

	if e != nil {
		return s.answerMessage(x, e, "cables not retrieved")
	}

	return response.SuccessAny(convert.Cables(result))
}
