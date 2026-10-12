package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/server"
)

func (s *Server) PutSecret(
	_ context.Context,
	r server.PutSecretRequestObject,
) (server.PutSecretResponseObject, error) {
	var values map[string]string

	if r.Body != nil {
		values = r.Body.Values
	}

	result, e := s.service.Merge(r.Params.Path, values)

	if validation.Is(e) {
		return server.PutSecret400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.PutSecret500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.PutSecret200JSONResponse{
		Path: r.Params.Path,
		Keys: result,
	}, nil
}
