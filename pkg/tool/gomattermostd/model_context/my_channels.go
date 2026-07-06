package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
	"sort"
)

func (s *Server) MyChannels(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.MyChannels,
) (*mcp.CallToolResult, error) {
	limit := a.Limit

	if limit <= 0 {
		limit = 30
	}

	me, e := s.client.Me()

	if e != nil {
		return s.captureDetail(e)
	}

	team := s.client.DefaultTeam()
	channels, e := s.client.Channels(team, me)

	if e != nil {
		return s.captureDetail(e)
	}

	sort.Slice(
		channels,
		func(a, b int) bool {
			return channels[a].LastPostAt > channels[b].LastPostAt
		},
	)

	if a.Since != "" {
		since, f := parseSince(a.Since)

		if f != nil {
			return response.Fail(
				"invalid since format: %v",
				f,
			)
		}

		sinceMilli := since.UnixMilli()
		var filtered []*model.Channel

		for _, c := range channels {
			if c.LastPostAt >= sinceMilli {
				filtered = append(filtered, c)
			}
		}

		channels = filtered
	}

	if a.TypeFilter != "" {
		var filtered []*model.Channel

		for _, c := range channels {
			if string(c.Type) == a.TypeFilter {
				filtered = append(filtered, c)
			}
		}

		channels = filtered
	}

	if len(channels) > limit {
		channels = channels[:limit]
	}

	type row struct {
		Identifier  string `json:"identifier"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Type        string `json:"type"`
		LastPostAt  string `json:"last_post_at"`
	}
	rows := make([]row, len(channels))

	for i, c := range channels {
		typeName := ""

		switch c.Type {
		case model.ChannelTypeOpen:
			typeName = "public"
		case model.ChannelTypePrivate:
			typeName = "private"
		case model.ChannelTypeDirect:
			typeName = "dm"
		case model.ChannelTypeGroup:
			typeName = "group_dm"
		}

		displayName := s.channelDisplayName(c)
		rows[i] = row{
			Identifier:  c.Id,
			Name:        c.Name,
			DisplayName: displayName,
			Type:        typeName,
			LastPostAt:  formatMilli(c.LastPostAt),
		}
	}

	return response.SuccessAny(rows)
}
