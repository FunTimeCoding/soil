package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) moveEntry(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Move,
) (*mcp.CallToolResult, error) {
	if e := s.service.Move(a.Identifier, a.Group); e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.Success("moved")
}
