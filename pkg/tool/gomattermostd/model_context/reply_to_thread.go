package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ReplyToThread(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ReplyToThread,
) (*mcp.CallToolResult, error) {
	if a.PostIdentifier == "" {
		return response.Fail("post_id is required")
	}

	if a.Message == "" {
		return response.Fail("message is required")
	}

	ch, e := s.service.Channel(a.ChannelIdentifier)

	if e != nil {
		return s.answerWith(e, "channel")
	}

	parent, e := s.service.FindPost(a.PostIdentifier)

	if e != nil {
		return s.answerWith(e, "post")
	}

	p, e := s.service.Reply(ch, parent, a.Message)

	if e != nil {
		return s.answer(e)
	}

	result := map[string]any{
		"id":         p.Id,
		"channel_id": p.ChannelId,
		"root_id":    p.RootId,
		"message":    p.Message,
		"create_at":  formatMilli(p.CreateAt),
	}

	if a.EmojiName != "" {
		result["emoji_name"] = a.EmojiName

		if f := s.service.React(parent, a.EmojiName); f != nil {
			result["reaction_error"] = f.Error()
		}
	}

	return response.SuccessAny(result)
}
