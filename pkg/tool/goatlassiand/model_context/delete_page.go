package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deletePage(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireString(generativeConstant.ParameterIdentifier)

	if f != nil {
		return response.Fail("identifier is required: %v", f)
	}

	draft := r.GetBool(constant.Draft, false)

	if g := s.service.DeletePage(identifier, draft); g != nil {
		return s.answer(g)
	}

	return response.Success("page deleted")
}
