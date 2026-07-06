package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CloseView(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CloseView,
) (*mcp.CallToolResult, error) {
	e := s.client.CloseView(int(a.ViewIdentifier))

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.Success("closed")
}
