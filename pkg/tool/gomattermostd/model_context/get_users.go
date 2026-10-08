package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	mattermostResponse "github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetUsers(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetUsers,
) (*mcp.CallToolResult, error) {
	limit := a.Limit

	if limit <= 0 {
		limit = 100
	}

	page := a.Page

	if page < 0 {
		page = 0
	}

	users, e := s.client.Users(page, limit)

	if e != nil {
		return s.captureDetail(e)
	}

	rows := make([]*mattermostResponse.User, len(users))

	for i, u := range users {
		rows[i] = mattermostResponse.NewUser(
			u.Id,
			u.Username,
			u.FirstName,
			u.LastName,
			u.Nickname,
			u.Email,
			u.IsBot,
		)
	}

	return response.SuccessAny(
		map[string]any{"users": rows, "page": page, "per_page": limit},
	)
}
