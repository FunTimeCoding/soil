package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/server"
)

func (s *Server) ShowSecret(
	x context.Context,
	r server.ShowSecretRequestObject,
) (server.ShowSecretResponseObject, error) {
	result, e := s.service.Keys(r.Params.Path)

	if validation.Is(e) {
		return server.ShowSecret400JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.ShowSecret404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.ShowSecret503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ShowSecret500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ShowSecret200JSONResponse{
		Path: r.Params.Path,
		Keys: result,
	}, nil
}
