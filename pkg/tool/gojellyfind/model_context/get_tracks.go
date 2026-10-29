package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetTracks(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetTracks,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Tracks(a.AlbumIdentifier)

	if e != nil {
		return s.answer(e)
	}

	return response.SuccessAny(result)
}
