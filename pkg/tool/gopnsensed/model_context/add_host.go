package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addHost(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, e := s.service.AddHost(
		hostRequest(r),
		r.GetBool(constant.ParameterApply, true),
	)

	if e != nil {
		return s.answer(e)
	}

	return response.Success(identifier)
}
