package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) PostRevocation(
	_ context.Context,
	r server.PostRevocationRequestObject,
) (server.PostRevocationResponseObject, error) {
	result, e := s.service.Revoke(r.Serial)

	if not_found.Is(e) {
		return server.PostRevocation404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PostRevocation500JSONResponse(*s.answer(e)), nil
	}

	return server.PostRevocation200JSONResponse(
		*convert.Certificate(result),
	), nil
}
