package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"
)

func (s *Server) GetThread(
	_ context.Context,
	r server.GetThreadRequestObject,
) (server.GetThreadResponseObject, error) {
	replies, e := s.service.ThreadOf(r.Identifier)

	if forbidden.Is(e) {
		return server.GetThread403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.GetThread404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetThread503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetThread500JSONResponse(*s.answer(e)), nil
	}

	return server.GetThread200JSONResponse(convertPosts(replies)), nil
}
