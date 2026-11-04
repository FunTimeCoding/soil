package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostRelease(
	x context.Context,
	r server.PostReleaseRequestObject,
) (server.PostReleaseResponseObject, error) {
	identifier, e := s.roster.ResolveByCallsign(r.Body.Callsign)

	if e != nil {
		return server.PostRelease500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	if e := s.roster.Release(identifier, r.Body.Callsign); e != nil {
		return server.PostRelease500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	return server.PostRelease200Response{}, nil
}
