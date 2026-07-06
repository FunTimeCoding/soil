package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) UpdateGroup(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.UpdateGroup,
) (*mcp.CallToolResult, error) {
	e := s.client.UpdateGroup(
		int(a.GroupIdentifier),
		a.Title,
		a.Color,
		nil,
	)

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.Success("group updated")
}
