package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) DeleteSession(
	_ context.Context,
	r server.DeleteSessionRequestObject,
) (server.DeleteSessionResponseObject, error) {
	if e := s.service.ReleaseByCallsign(r.Callsign); e != nil {
		return server.DeleteSession500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.DeleteSession200Response{}, nil
}
