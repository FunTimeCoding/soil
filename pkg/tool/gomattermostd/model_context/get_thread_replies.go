package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetThreadReplies(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetThreadReplies,
) (*mcp.CallToolResult, error) {
	if a.PostID == "" {
		return response.Fail("post_id is required")
	}

	parent, e := s.client.FindPost(a.PostID)

	if e != nil {
		return s.captureFail(e, "post not found")
	}

	replies, f := s.client.Thread(parent)

	if f != nil {
		return s.captureDetail(f)
	}

	type row struct {
		ID       string   `json:"id"`
		Username string   `json:"username"`
		Message  string   `json:"message"`
		CreateAt string   `json:"create_at"`
		FileIds  []string `json:"file_ids,omitempty"`
	}
	rows := make([]row, len(replies))

	for i, r := range replies {
		rows[i] = row{
			ID:       r.Raw.Id,
			Message:  r.Message,
			CreateAt: formatTime(r.Create),
		}

		if r.User != nil {
			rows[i].Username = r.User.Username
		}

		if len(r.Raw.FileIds) > 0 {
			rows[i].FileIds = r.Raw.FileIds
		}
	}

	return response.SuccessAny(
		map[string]any{
			"post_id": a.PostID,
			"replies": rows,
		},
	)
}
