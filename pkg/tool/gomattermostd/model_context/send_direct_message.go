package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SendDirectMessage(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SendDirectMessage,
) (*mcp.CallToolResult, error) {
	if a.Username == "" {
		return response.Fail("username is required")
	}

	if a.Message == "" {
		return response.Fail("message is required")
	}

	user, resp, e := s.client.Nested().GetUserByUsername(
		s.client.Context(),
		a.Username,
		"",
	)

	if e != nil {
		if resp != nil && resp.StatusCode == 404 {
			return response.Fail("user not found: %s", a.Username)
		}

		return s.captureDetail(e)
	}

	p, e := s.client.DirectMessage(user, a.Message)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(
		map[string]any{
			"id":         p.Id,
			"channel_id": p.ChannelId,
			"to":         a.Username,
			"message":    p.Message,
			"create_at":  formatMilli(p.CreateAt),
		},
	)
}
