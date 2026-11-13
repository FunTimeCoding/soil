package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getRun(
	x context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier := uint(r.GetInt(constant.Identifier, 0))

	if identifier == 0 {
		return response.Fail("id is required")
	}

	result, e := s.service.Find(identifier)

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(result)
}
