package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostAnnounce(
	_ context.Context,
	r server.PostAnnounceRequestObject,
) (server.PostAnnounceResponseObject, error) {
	var files string

	if r.Body.Files != nil {
		files = join.NewLine(*r.Body.Files)
	}

	identifier, e := s.service.ResolveByCallsign(r.Body.Callsign)

	if e != nil {
		return server.PostAnnounce500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	if e := s.service.Announce(identifier, r.Body.Callsign, r.Body.Topic, files); e != nil {
		if not_found.Is(e) {
			return server.PostAnnounce404JSONResponse(
				server.Error{Error: e.Error()},
			), nil
		}

		return server.PostAnnounce500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostAnnounce200Response{}, nil
}
