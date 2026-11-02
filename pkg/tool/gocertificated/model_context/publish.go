package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) publish(
	x context.Context,
	_ mcp.CallToolRequest,
	_ struct{},
) (*mcp.CallToolResult, error) {
	commit, change, e := s.service.Publish(x)

	if e != nil {
		return s.answer(e)
	}

	if commit == "" {
		return response.Success("Nothing to publish")
	}

	return response.SuccessAny(
		map[string]any{
			"commit":    commit,
			"published": convert.Changes(change),
		},
	)
}
