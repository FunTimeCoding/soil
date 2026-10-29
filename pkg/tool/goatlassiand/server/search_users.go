package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) SearchUsers(
	_ context.Context,
	r server.SearchUsersRequestObject,
) (server.SearchUsersResponseObject, error) {
	users, e := s.service.SearchUsers(r.Params.Query)

	if forbidden.Is(e) {
		return server.SearchUsers403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.SearchUsers503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.SearchUsers500JSONResponse(*s.answer(e)), nil
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
