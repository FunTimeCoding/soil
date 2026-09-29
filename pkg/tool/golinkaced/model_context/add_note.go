package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) AddNote(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.AddNote,
) (*mcp.CallToolResult, error) {
	result, e := s.client.CreateNote(a.LinkIdentifier, a.Text)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(
		map[string]any{
			"identifier": result.Identifier,
			"text":       result.Text,
		},
	)
}
