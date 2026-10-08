package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	mattermostResponse "github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/response"
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

	var rows []*mattermostResponse.UserMatch

	for _, u := range autocomplete.Users {
		rows = append(
			rows,
			mattermostResponse.NewUserMatch(
				u.Id,
				u.Username,
				u.FirstName,
				u.LastName,
				u.Nickname,
				u.Email,
			),
		)
	}

	return response.SuccessAny(rows)
}
