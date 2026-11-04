package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) profile(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, _, e := s.service.Profile(
		q.GetString(constant.Topic, ""),
		q.GetString(constant.Scope, ""),
		false,
	)

	if e != nil {
		return s.answer(x, e)
	}

	return response.Success(result.Text)
}
