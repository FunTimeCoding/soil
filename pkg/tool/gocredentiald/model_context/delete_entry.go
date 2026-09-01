package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteEntry(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Identifier,
) (*mcp.CallToolResult, error) {
	if e := s.service.Delete(a.Identifier); e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.Success("deleted")
}
