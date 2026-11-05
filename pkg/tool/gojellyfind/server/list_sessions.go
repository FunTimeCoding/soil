package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func (s *Server) ListSessions(
	x context.Context,
	_ server.ListSessionsRequestObject,
) (server.ListSessionsResponseObject, error) {
	sessions, e := s.service.Sessions()

	if unavailable.Is(e) {
		return server.ListSessions503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListSessions500JSONResponse(*s.answer(x, e)), nil
	}

	var result server.ListSessions200JSONResponse

	for _, v := range sessions {
		result = append(result, convertSession(v))
	}

	return result, nil
}
