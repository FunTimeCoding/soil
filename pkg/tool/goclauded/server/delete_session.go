package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) DeleteSession(
	x context.Context,
	r server.DeleteSessionRequestObject,
) (server.DeleteSessionResponseObject, error) {
	if e := s.roster.ReleaseByCallsign(r.Callsign); e != nil {
		return server.DeleteSession500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	return server.DeleteSession200Response{}, nil
}
