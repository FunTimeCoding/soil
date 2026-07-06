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
	if a.PostID == "" {
		return response.Fail("post_id is required")
	}

	if a.Message == "" {
		return response.Fail("message is required")
	}

	ch, e := s.client.Channel(a.ChannelID)

	if e != nil {
		return s.captureFail(e, "channel not found")
	}

	parent, e := s.client.FindPost(a.PostID)

	if e != nil {
		return s.captureFail(e, "post not found")
	}

	p, e := s.client.Reply(ch, parent, a.Message)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(
		map[string]any{
			"id":         p.Id,
			"channel_id": p.ChannelId,
			"root_id":    p.RootId,
			"message":    p.Message,
			"create_at":  formatMilli(p.CreateAt),
		},
	)
}
