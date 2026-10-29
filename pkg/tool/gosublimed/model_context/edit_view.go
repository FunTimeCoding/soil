package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/types/edit"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) EditView(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.EditView,
) (*mcp.CallToolResult, error) {
	v, e := s.service.EditView(
		edit.New(
			int(a.ViewIdentifier),
			a.OldString,
			a.NewString,
			bool(a.ReplaceAll),
		),
	)

	if e != nil {
		return s.answer(e)
	}

	return response.SuccessAny(v)
}
