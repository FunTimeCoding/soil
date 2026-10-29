package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
)

func (s *Server) PostUser(
	_ context.Context,
	r server.PostUserRequestObject,
) (server.PostUserResponseObject, error) {
	result, e := s.service.CreateUser(
		r.Body.Account,
		r.Body.Name,
		r.Body.Surname,
		text(r.Body.Mail),
		text(r.Body.Password),
	)

	if validation.Is(e) {
		return server.PostUser400JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.PostUser403JSONResponse{Error: e.Error()}, nil
	}

	if conflict.Is(e) {
		return server.PostUser409JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.PostUser503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PostUser500JSONResponse(*s.answer(e)), nil
	}

	return server.PostUser200JSONResponse(*convert.User(result)), nil
}
