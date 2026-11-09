package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostAnnounce(
	x context.Context,
	r server.PostAnnounceRequestObject,
) (server.PostAnnounceResponseObject, error) {
	var files string

	if r.Body.Files != nil {
		files = join.NewLine(*r.Body.Files)
	}

	identifier, e := s.roster.ResolveCallsign(r.Body.Callsign)

	if e != nil {
		return server.PostAnnounce500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	if e := s.roster.Announce(
		identifier,
		r.Body.Callsign,
		r.Body.Topic,
		files,
	); e != nil {
		if not_found.Is(e) {
			return server.PostAnnounce404JSONResponse(
				server.Error{Error: e.Error()},
			), nil
		}

		return server.PostAnnounce500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	return server.PostAnnounce200Response{}, nil
}
