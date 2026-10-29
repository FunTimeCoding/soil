package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SaveView(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SaveView,
) (*mcp.CallToolResult, error) {
	if e := s.service.SaveView(int(a.ViewIdentifier), a.FilePath); e != nil {
		return s.answer(e)
	}

	return response.Success("saved")
}
