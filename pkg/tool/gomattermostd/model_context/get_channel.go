package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Server) GetChannel(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetChannel,
) (*mcp.CallToolResult, error) {
	if a.ChannelID == "" && a.ChannelName == "" {
		return response.Fail("channel_id or channel_name is required")
	}

	var ch *model.Channel

	if a.ChannelName != "" {
		var e error
		ch, e = s.client.TeamChannel(a.ChannelName)

		if e != nil {
			return s.captureFail(e, "channel not found")
		}
	} else {
		var e error
		ch, e = s.client.Channel(a.ChannelID)

		if e != nil {
			return s.captureFail(e, "channel not found")
		}
	}

	result := map[string]any{
		"id":           ch.Id,
		"name":         ch.Name,
		"display_name": s.channelDisplayName(ch),
		"type":         channelTypeName(ch.Type),
		"header":       ch.Header,
		"purpose":      ch.Purpose,
	}

	if ch.LastPostAt > 0 {
		result["last_post_at"] = formatMilli(ch.LastPostAt)
	}

	return response.SuccessAny(result)
}
