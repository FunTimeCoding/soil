package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteHost(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier := r.GetString(generativeConstant.ParameterIdentifier, "")

	if e := s.service.DeleteHost(
		identifier,
		r.GetBool(constant.ParameterApply, true),
	); e != nil {
		return s.answer(e)
	}

	return response.Success(identifier)
}
