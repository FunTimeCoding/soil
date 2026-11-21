package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListSystems(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListSystems,
) (*mcp.CallToolResult, error) {
	result, e := s.system.Summaries(a.Score)

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(
		map[string]any{"systems": convert.SystemSummaries(result)},
	)
}
