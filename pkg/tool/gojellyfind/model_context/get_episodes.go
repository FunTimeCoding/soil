package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetEpisodes(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetEpisodes,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Episodes(a.SeriesIdentifier)

	if e != nil {
		return s.answer(e)
	}

	return response.SuccessAny(result)
}
