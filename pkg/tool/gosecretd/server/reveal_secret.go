package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/server"
)

func (s *Server) RevealSecret(
	_ context.Context,
	r server.RevealSecretRequestObject,
) (server.RevealSecretResponseObject, error) {
	result, e := s.service.Reveal(r.Params.Path, r.Params.Key)

	if validation.Is(e) {
		return server.RevealSecret400JSONResponse(*clientError(e)), nil
	}

	if not_found.Is(e) {
		return server.RevealSecret404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.RevealSecret500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.RevealSecret200JSONResponse{
		Path:  r.Params.Path,
		Key:   r.Params.Key,
		Value: result,
	}, nil
}
