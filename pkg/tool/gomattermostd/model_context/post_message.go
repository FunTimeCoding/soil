package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Server) PostMessage(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.PostMessage,
) (*mcp.CallToolResult, error) {
	if a.Message == "" {
		return response.Fail("message is required")
	}

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

	p, e := s.client.PostSimple(ch, a.Message)

	if e != nil {
		return s.captureDetail(e)
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
