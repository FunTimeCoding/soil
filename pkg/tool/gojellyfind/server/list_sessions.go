package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func (s *Server) ListSessions(
	_ context.Context,
	_ server.ListSessionsRequestObject,
) (server.ListSessionsResponseObject, error) {
	sessions, e := s.client.Sessions()

	if e != nil {
		return server.ListSessions500JSONResponse(*s.captureDetail(e)), nil
	}

	var result server.ListSessions200JSONResponse

	for _, v := range sessions {
		result = append(result, convertSession(v))
	}

	return result, nil
}
