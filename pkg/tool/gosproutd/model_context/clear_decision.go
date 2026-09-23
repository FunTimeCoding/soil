package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) clearDecision(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ClearDecision,
) (*mcp.CallToolResult, error) {
	if a.Understanding == "" {
		return response.Fail(
			"clearing a decision records what you understood - say it in one line",
		)
	}

	s.service.Clear(uint(a.Identifier), a.Understanding, a.Heard)

	return response.Success("cleared")
}
