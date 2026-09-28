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
	if a.AlbumIdentifier == "" {
		return response.Fail("album_id is required")
	}

	result, e := s.client.Tracks(a.AlbumIdentifier)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
