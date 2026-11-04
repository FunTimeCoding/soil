package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) AddNote(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.AddNote,
) (*mcp.CallToolResult, error) {
	result, e := s.service.CreateNote(a.LinkIdentifier, a.Text)

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(
		map[string]any{
			"identifier": result.Identifier,
			"text":       result.Text,
		},
	)
}
