package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostRelease(
	_ context.Context,
	r server.PostReleaseRequestObject,
) (server.PostReleaseResponseObject, error) {
	identifier, e := s.service.ResolveByCallsign(r.Body.Callsign)

	if e != nil {
		return server.PostRelease500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	if e := s.service.Release(identifier, r.Body.Callsign); e != nil {
		return server.PostRelease500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostRelease200Response{}, nil
}
