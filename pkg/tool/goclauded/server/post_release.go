package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostRelease(
	x context.Context,
	r server.PostReleaseRequestObject,
) (server.PostReleaseResponseObject, error) {
	identifier, e := s.roster.ResolveCallsign(r.Body.Callsign)

	if e != nil {
		return server.PostRelease500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	if identifier == "" {
		return server.PostRelease404JSONResponse(
			server.Error{
				Error: not_found.New(
					constant.SessionKind,
					r.Body.Callsign,
				).Error(),
			},
		), nil
	}

	if e := s.roster.Release(identifier, r.Body.Callsign); e != nil {
		return server.PostRelease500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	return server.PostRelease200Response{}, nil
}
