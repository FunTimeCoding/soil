package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
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

	type row struct {
		ID        string `json:"id"`
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Nickname  string `json:"nickname"`
		Email     string `json:"email"`
		IsBot     bool   `json:"is_bot"`
	}
	rows := make([]row, len(users))

	for i, u := range users {
		rows[i] = row{
			ID:        u.Id,
			Username:  u.Username,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Nickname:  u.Nickname,
			Email:     u.Email,
			IsBot:     u.IsBot,
		}
	}

	return response.SuccessAny(
		map[string]any{
			"users":    rows,
			"page":     page,
			"per_page": limit,
		},
	)
}
