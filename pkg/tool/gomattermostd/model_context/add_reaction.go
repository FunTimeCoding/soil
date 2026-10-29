package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) AddReaction(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.AddReaction,
) (*mcp.CallToolResult, error) {
	if a.PostIdentifier == "" {
		return response.Fail("post_id is required")
	}

	if a.EmojiName == "" {
		return response.Fail("emoji_name is required")
	}

	p, e := s.service.FindPost(a.PostIdentifier)

	if e != nil {
		return s.answer(e)
	}

	e = s.service.React(p, a.EmojiName)

	if e != nil {
		return s.answer(e)
	}

	return response.SuccessAny(
		map[string]any{"post_id": a.PostIdentifier, "emoji_name": a.EmojiName},
	)
}
