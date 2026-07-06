package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) EditView(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.EditView,
) (*mcp.CallToolResult, error) {
	if a.OldString == "" {
		return response.Fail("old_string is required")
	}

	v, e := s.client.EditView(
		int(a.ViewIdentifier),
		a.OldString,
		a.NewString,
		bool(a.ReplaceAll),
	)

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.SuccessAny(v)
}
