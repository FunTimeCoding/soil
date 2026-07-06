package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ReadTab(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ReadTab,
) (*mcp.CallToolResult, error) {
	v, e := s.client.ReadTab(
		int(a.TabIdentifier),
		bool(a.Raw),
	)

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.SuccessAny(v)
}
