package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	mattermostResponse "github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetThreadReplies(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetThreadReplies,
) (*mcp.CallToolResult, error) {
	if a.PostIdentifier == "" {
		return response.Fail("post_id is required")
	}

	parent, e := s.client.FindPost(a.PostIdentifier)

	if e != nil {
		return s.captureFail(e, "post not found")
	}

	replies, f := s.client.Thread(parent)

	if f != nil {
		return s.captureDetail(f)
	}

	rows := make([]*mattermostResponse.Reply, len(replies))

	for i, r := range replies {
		rows[i] = mattermostResponse.NewReply(
			r.Raw.Id,
			r.Message,
			formatTime(r.Create),
		)

		if r.User != nil {
			rows[i].Username = r.User.Username
		}

		if len(r.Raw.FileIds) > 0 {
			rows[i].FileIds = r.Raw.FileIds
		}
	}

	return response.SuccessAny(
		map[string]any{"post_id": a.PostIdentifier, "replies": rows},
	)
}
