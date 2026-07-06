package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
	"strings"
)

func (s *Server) SearchChannels(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SearchChannels,
) (*mcp.CallToolResult, error) {
	if a.Query == "" {
		return response.Fail("query is required")
	}

	me, e := s.client.Me()

	if e != nil {
		return s.captureDetail(e)
	}

	team := s.client.DefaultTeam()
	channels, _, e := s.client.Nested().AutocompleteChannelsForTeamForSearch(
		s.client.Context(),
		team.Id,
		a.Query,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	type row struct {
		Identifier  string `json:"identifier"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Type        string `json:"type"`
		LastPostAt  string `json:"last_post_at,omitempty"`
	}
	var rows []row

	for _, c := range channels {
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

		displayName := c.DisplayName

		if c.Type == model.ChannelTypeDirect && displayName == "" {
			parts := strings.Split(c.Name, "__")

			for _, p := range parts {
				if p != me.Id {
					other, f := s.client.User(p)

					if f == nil && other != nil {
						displayName = other.Username
					}

					break
				}
			}
		}

		r := row{
			Identifier:  c.Id,
			Name:        c.Name,
			DisplayName: displayName,
			Type:        typeName,
		}

		if c.LastPostAt > 0 {
			r.LastPostAt = formatMilli(c.LastPostAt)
		}

		rows = append(rows, r)
	}

	return response.SuccessAny(rows)
}
