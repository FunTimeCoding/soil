package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	mattermostResponse "github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SearchMessages(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SearchMessages,
) (*mcp.CallToolResult, error) {
	if a.Terms == "" {
		return response.Fail("terms is required")
	}

	team := s.client.DefaultTeam()
	list, _, e := s.client.Nested().SearchPosts(
		s.client.Context(),
		team.Id,
		a.Terms,
		false,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	posts := post.NewSlice(post.FromList(list, true))
	g := s.client.Enrich(posts)

	if g != nil {
		return s.captureDetail(g)
	}

	var rows []*mattermostResponse.MessageMatch

	for _, p := range posts {
		r := mattermostResponse.NewMessageMatch(
			p.Raw.Id,
			p.Message,
			formatTime(p.Create),
		)

		if p.User != nil {
			r.Username = p.User.Username
		}

		if p.Raw.ChannelId != "" {
			ch, f := s.client.Channel(p.Raw.ChannelId)

			if f == nil {
				r.Channel = s.channelDisplayName(ch)
			}
		}

		if len(p.Raw.FileIds) > 0 {
			r.FileIds = p.Raw.FileIds
		}

		rows = append(rows, r)
	}

	return response.SuccessAny(
		map[string]any{"results": rows, "count": len(rows)},
	)
}
