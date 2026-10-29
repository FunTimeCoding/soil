package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) PostPublication(
	_ context.Context,
	_ server.PostPublicationRequestObject,
) (server.PostPublicationResponseObject, error) {
	commit, change, e := s.service.Publish()

	if unavailable.Is(e) {
		return server.PostPublication503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PostPublication500JSONResponse(*s.answer(e)), nil
	}

	result := server.PublicationResponse{Pending: convert.Changes(change)}

	if commit != "" {
		result.Commit = &commit
	}

	return server.PostPublication200JSONResponse(result), nil
}
