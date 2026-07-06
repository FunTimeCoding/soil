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
	if a.PostID == "" {
		return response.Fail("post_id is required")
	}

	if a.EmojiName == "" {
		return response.Fail("emoji_name is required")
	}

	p, e := s.client.FindPost(a.PostID)

	if e != nil {
		return s.captureFail(e, "post not found")
	}

	e = s.client.React(p, a.EmojiName)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(
		map[string]any{
			"post_id":    a.PostID,
			"emoji_name": a.EmojiName,
		},
	)
}
