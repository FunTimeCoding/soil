package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listRules(
	x context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Rules(
		r.GetString(generativeConstant.ParameterQuery, ""),
	)

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(convert.Rules(result))
}
