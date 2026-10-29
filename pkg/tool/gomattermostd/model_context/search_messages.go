package model_context

import (
	"context"
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

	posts, e := s.service.SearchPosts(a.Terms)

	if e != nil {
		return s.answer(e)
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
			ch, f := s.service.Channel(p.Raw.ChannelId)

			if f == nil {
				r.Channel = s.service.ChannelDisplayName(ch)
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
