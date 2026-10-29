package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/server"
)

func (s *Server) DeleteClaim(
	_ context.Context,
	r server.DeleteClaimRequestObject,
) (server.DeleteClaimResponseObject, error) {
	e := s.service.Release(r.Params.Item, r.Params.Owner)

	if validation.Is(e) {
		return server.DeleteClaim400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteClaim500JSONResponse(*s.answer(e)), nil
	}

	return server.DeleteClaim204Response{}, nil
}
