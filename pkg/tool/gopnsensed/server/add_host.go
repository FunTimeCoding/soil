package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) AddHost(
	_ context.Context,
	r server.AddHostRequestObject,
) (server.AddHostResponseObject, error) {
	identifier, e := s.service.AddHost(
		convert.HostRequest(r.Body),
		apply(r.Params.Apply),
	)

	if validation.Is(e) {
		return server.AddHost400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.AddHost500JSONResponse(*s.answer(e)), nil
	}

	return server.AddHost200JSONResponse(*convert.Identifier(identifier)), nil
}
