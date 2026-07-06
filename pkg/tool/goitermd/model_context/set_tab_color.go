package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SetTabColor(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SetTabColor,
) (*mcp.CallToolResult, error) {
	e := s.client.SetTabColor(
		a.SessionIdentifier,
		int(a.Red),
		int(a.Green),
		int(a.Blue),
	)

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.Success("color set")
}
