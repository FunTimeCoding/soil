package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	mattermostResponse "github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/response"
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

	me, e := s.service.Me()

	if e != nil {
		return s.answer(e)
	}

	channels, e := s.service.SearchChannels(a.Query)

	if e != nil {
		return s.answer(e)
	}

	var rows []*mattermostResponse.ChannelMatch

	for _, c := range channels {
		typeName := channelTypeName(c.Type)
		displayName := c.DisplayName

		if c.Type == model.ChannelTypeDirect && displayName == "" {
			parts := strings.Split(c.Name, "__")

			for _, p := range parts {
				if p != me.Id {
					other, f := s.service.User(p)

					if f == nil && other != nil {
						displayName = other.Username
					}

					break
				}
			}
		}

		r := mattermostResponse.NewChannelMatch(
			c.Id,
			c.Name,
			displayName,
			typeName,
		)

		if c.LastPostAt > 0 {
			r.LastPostAt = formatMilli(c.LastPostAt)
		}

		rows = append(rows, r)
	}

	return response.SuccessAny(rows)
}
