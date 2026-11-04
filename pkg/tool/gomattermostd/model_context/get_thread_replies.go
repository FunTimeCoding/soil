package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	mattermostResponse "github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetThreadReplies(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.GetThreadReplies,
) (*mcp.CallToolResult, error) {
	if a.PostIdentifier == "" {
		return response.Fail("post_id is required")
	}

	parent, e := s.service.FindPost(a.PostIdentifier)

	if e != nil {
		return s.answer(x, e)
	}

	replies, f := s.service.Thread(parent)

	if f != nil {
		return s.answer(x, f)
	}

	rows := make([]*mattermostResponse.Reply, len(replies))

	for i, r := range replies {
		rows[i] = replyRow(r)
	}

	return response.SuccessAny(
		map[string]any{"post_id": a.PostIdentifier, "replies": rows},
	)
}
