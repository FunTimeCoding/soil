package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SearchUsers(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SearchUsers,
) (*mcp.CallToolResult, error) {
	if a.Query == "" {
		return response.Fail("query is required")
	}

	team := s.client.DefaultTeam()
	autocomplete, _, e := s.client.Nested().AutocompleteUsersInTeam(
		s.client.Context(),
		team.Id,
		a.Query,
		20,
		"",
	)

	if e != nil {
		return s.captureDetail(e)
	}

	type row struct {
		Identifier string `json:"identifier"`
		Username   string `json:"username"`
		FirstName  string `json:"first_name,omitempty"`
		LastName   string `json:"last_name,omitempty"`
		Nickname   string `json:"nickname,omitempty"`
		Email      string `json:"email,omitempty"`
	}
	var rows []row

	for _, u := range autocomplete.Users {
		rows = append(
			rows,
			row{
				Identifier: u.Id,
				Username:   u.Username,
				FirstName:  u.FirstName,
				LastName:   u.LastName,
				Nickname:   u.Nickname,
				Email:      u.Email,
			},
		)
	}

	return response.SuccessAny(rows)
}
