package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
)

func (s *Server) PostUserAccountPassword(
	_ context.Context,
	r server.PostUserAccountPasswordRequestObject,
) (server.PostUserAccountPasswordResponseObject, error) {
	e := s.service.SetPassword(r.Account, r.Body.Password)

	if validation.Is(e) {
		return server.PostUserAccountPassword400JSONResponse{
			Error: e.Error(),
		}, nil
	}

	if not_found.Is(e) {
		return server.PostUserAccountPassword404JSONResponse{
			Error: e.Error(),
		}, nil
	}

	if forbidden.Is(e) {
		return server.PostUserAccountPassword403JSONResponse{
			Error: e.Error(),
		}, nil
	}

	if unavailable.Is(e) {
		return server.PostUserAccountPassword503JSONResponse{
			Error: e.Error(),
		}, nil
	}

	if e != nil {
		return server.PostUserAccountPassword500JSONResponse(*s.answer(e)), nil
	}

	return server.PostUserAccountPassword200JSONResponse{
		Message: fmt.Sprintf("password set: %s", r.Account),
	}, nil
}
