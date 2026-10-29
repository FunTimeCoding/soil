package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) PostSigningRequest(
	_ context.Context,
	r server.PostSigningRequestRequestObject,
) (server.PostSigningRequestResponseObject, error) {
	result, e := s.service.SignRequest(r.Body)

	if not_found.Is(e) {
		return server.PostSigningRequest404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PostSigningRequest500JSONResponse(*s.answer(e)), nil
	}

	return server.PostSigningRequest200JSONResponse(
		*convert.Certificate(result),
	), nil
}
