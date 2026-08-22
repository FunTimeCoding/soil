package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) SearchUsers(
	x context.Context,
	r server.SearchUsersRequestObject,
) (server.SearchUsersResponseObject, error) {
	users, _, e := s.jira.Nested().User.FindWithContext(x, r.Params.Query)

	if e != nil {
		return server.SearchUsers500JSONResponse(*s.captureDetail(e)), nil
	}

	var result server.SearchUsers200JSONResponse

	for _, u := range convert.JiraUsers(users) {
		row := &server.JiraUser{
			AccountIdentifier: u.AccountIdentifier,
			DisplayName:       u.DisplayName,
			Active:            u.Active,
		}

		if u.Email != "" {
			row.Email = new(u.Email)
		}

		result = append(result, row)
	}

	return result, nil
}
