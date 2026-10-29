package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) PostMessage(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.PostMessage,
) (*mcp.CallToolResult, error) {
	if a.Message == "" {
		return response.Fail("message is required")
	}

	if a.ChannelIdentifier == "" && a.ChannelName == "" {
		return response.Fail("channel_id or channel_name is required")
	}

	ch, e := s.service.ResolveChannel(a.ChannelIdentifier, a.ChannelName)

	if e != nil {
		return s.answer(e)
	}

	p, e := s.service.PostSimple(ch, a.Message)

	if e != nil {
		return s.answer(e)
	}

	return response.SuccessAny(
		map[string]any{
			"id":         p.Id,
			"channel_id": p.ChannelId,
			"message":    p.Message,
			"create_at":  formatMilli(p.CreateAt),
		},
	)
}
