package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Server) MyThreads(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.MyThreads,
) (*mcp.CallToolResult, error) {
	limit := a.Limit

	if limit <= 0 {
		limit = 20
	}

	me, e := s.client.Me()

	if e != nil {
		return s.captureDetail(e)
	}

	team := s.client.DefaultTeam()
	opts := model.GetUserThreadsOpts{
		PageSize: uint64(limit),
		Extended: true,
	}

	if a.Since != "" {
		since, g := parseSince(a.Since)

		if g != nil {
			return response.Fail(
				"invalid since format: %v",
				g,
			)
		}

		opts.Since = uint64(since.UnixMilli())
	}

	threads, _, e := s.client.Nested().GetUserThreads(
		s.client.Context(),
		me.Id,
		team.Id,
		opts,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	type participant struct {
		Username string `json:"username"`
	}
	type row struct {
		PostIdentifier string        `json:"post_identifier"`
		Channel        string        `json:"channel,omitempty"`
		Author         string        `json:"author"`
		Message        string        `json:"message"`
		ReplyCount     int64         `json:"reply_count"`
		UnreadReplies  int64         `json:"unread_replies"`
		UnreadMentions int64         `json:"unread_mentions"`
		LastReplyAt    string        `json:"last_reply_at"`
		Participants   []participant `json:"participants,omitempty"`
	}
	var rows []row

	for _, thread := range threads.Threads {
		if a.UnreadOnly && thread.UnreadReplies == 0 {
			continue
		}

		r := row{
			PostIdentifier: thread.PostId,
			ReplyCount:     thread.ReplyCount,
			UnreadReplies:  thread.UnreadReplies,
			UnreadMentions: thread.UnreadMentions,
			LastReplyAt:    formatMilli(thread.LastReplyAt),
		}

		if thread.Post != nil {
			r.Message = thread.Post.Message

			if thread.Post.UserId != "" {
				author, f := s.client.User(thread.Post.UserId)

				if f == nil {
					r.Author = author.Username
				}
			}

			ch, f := s.client.Channel(thread.Post.ChannelId)

			if f == nil {
				r.Channel = s.channelDisplayName(ch)
			}
		}

		for _, p := range thread.Participants {
			if p != nil {
				r.Participants = append(
					r.Participants,
					participant{Username: p.Username},
				)
			}
		}

		rows = append(rows, r)
	}

	return response.SuccessAny(
		map[string]any{
			"threads":               rows,
			"total":                 threads.Total,
			"total_unread_threads":  threads.TotalUnreadThreads,
			"total_unread_mentions": threads.TotalUnreadMentions,
		},
	)
}
