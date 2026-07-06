package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SetTabTitle(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SetTabTitle,
) (*mcp.CallToolResult, error) {
	if a.Title == "" {
		return response.Fail("title is required")
	}

	e := s.client.SetTabTitle(a.TabIdentifier, a.Title)

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.Success("set title: %s", a.Title)
}
