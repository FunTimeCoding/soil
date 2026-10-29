package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetChannel(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetChannel,
) (*mcp.CallToolResult, error) {
	if a.ChannelIdentifier == "" && a.ChannelName == "" {
		return response.Fail("channel_id or channel_name is required")
	}

	ch, e := s.service.ResolveChannel(a.ChannelIdentifier, a.ChannelName)

	if e != nil {
		return s.answer(e)
	}

	result := map[string]any{
		"id":           ch.Id,
		"name":         ch.Name,
		"display_name": s.service.ChannelDisplayName(ch),
		"type":         channelTypeName(ch.Type),
		"header":       ch.Header,
		"purpose":      ch.Purpose,
	}

	if ch.LastPostAt > 0 {
		result["last_post_at"] = formatMilli(ch.LastPostAt)
	}

	return response.SuccessAny(result)
}
