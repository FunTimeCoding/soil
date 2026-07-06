package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Server) GetChannelHistory(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetChannelHistory,
) (*mcp.CallToolResult, error) {
	if a.ChannelID == "" && a.ChannelName == "" {
		return response.Fail("channel_id or channel_name is required")
	}

	limit := a.Limit

	if limit <= 0 {
		limit = 30
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

	var posts []*post.Post
	var g error
	truncated := 0

	if a.Since != "" {
		since, f := parseSince(a.Since)

		if f != nil {
			return response.Fail(
				"invalid since format: %v",
				f,
			)
		}

		posts, g = s.client.RecentPosts(ch, since.UnixMilli())

		if g != nil {
			return s.captureDetail(g)
		}

		if len(posts) > limit {
			truncated = len(posts) - limit
			posts = posts[truncated:]
		}
	} else {
		posts, g = s.client.LatestPosts(ch, limit)

		if g != nil {
			return s.captureDetail(g)
		}
	}

	type row struct {
		ID       string   `json:"id"`
		Username string   `json:"username"`
		Message  string   `json:"message"`
		CreateAt string   `json:"create_at"`
		RootID   *string  `json:"root_id,omitempty"`
		FileIds  []string `json:"file_ids,omitempty"`
	}
	rows := make([]row, len(posts))

	for i, p := range posts {
		r := row{
			ID:       p.Raw.Id,
			Message:  p.Message,
			CreateAt: formatTime(p.Create),
		}

		if p.User != nil {
			r.Username = p.User.Username
		}

		if p.Raw.RootId != "" {
			r.RootID = new(p.Raw.RootId)
		}

		if len(p.Raw.FileIds) > 0 {
			r.FileIds = p.Raw.FileIds
		}

		rows[i] = r
	}

	result := map[string]any{
		"posts": rows,
		"limit": limit,
	}

	if truncated > 0 {
		result["note"] = fmt.Sprintf(
			"showing the newest %d of %d posts since %s - narrow since or raise limit for more",
			limit,
			limit+truncated,
			a.Since,
		)
	}

	return response.SuccessAny(result)
}
