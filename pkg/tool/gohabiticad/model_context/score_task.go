package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) scoreTask(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireString(generativeConstant.ParameterIdentifier)

	if f != nil {
		return response.Fail("identifier is required: %v", f)
	}

	direction := r.GetString(constant.Direction, "up")
	result, g := s.habitica.Score(identifier, direction)

	if g != nil {
		return s.answer(g)
	}

	return response.SuccessAny(convert.ScoreResult(result))
}
