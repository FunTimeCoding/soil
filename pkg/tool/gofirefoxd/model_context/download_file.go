package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) DownloadFile(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DownloadFile,
) (*mcp.CallToolResult, error) {
	v, e := s.service.DownloadFile(a.Locator, a.Path)

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(v)
}
