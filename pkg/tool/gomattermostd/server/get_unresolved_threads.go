package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"
)

func (s *Server) GetUnresolvedThreads(
	_ context.Context,
	r server.GetUnresolvedThreadsRequestObject,
) (server.GetUnresolvedThreadsResponseObject, error) {
	threads, e := s.service.UnresolvedThreads(r.Name, r.Params.Since)

	if forbidden.Is(e) {
		return server.GetUnresolvedThreads403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.GetUnresolvedThreads404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetUnresolvedThreads503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetUnresolvedThreads500JSONResponse(*s.answer(e)), nil
	}

	return server.GetUnresolvedThreads200JSONResponse(
		convertThreads(threads),
	), nil
}
