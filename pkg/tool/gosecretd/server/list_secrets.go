package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/server"
)

func (s *Server) ListSecrets(
	_ context.Context,
	r server.ListSecretsRequestObject,
) (server.ListSecretsResponseObject, error) {
	path := ""

	if r.Params.Path != nil {
		path = *r.Params.Path
	}

	result, e := s.service.List(path)

	if validation.Is(e) {
		return server.ListSecrets400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListSecrets500JSONResponse(*s.answer(e)), nil
	}

	return server.ListSecrets200JSONResponse{Path: path, Entries: result}, nil
}
