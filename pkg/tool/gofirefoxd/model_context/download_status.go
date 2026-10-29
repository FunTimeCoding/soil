package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) DownloadStatus(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.DownloadStatus,
) (*mcp.CallToolResult, error) {
	v, e := s.service.DownloadStatus(int(a.DownloadIdentifier))

	if e != nil {
		return s.answer(e)
	}

	return response.SuccessAny(v)
}
