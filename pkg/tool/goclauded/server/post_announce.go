package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
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

	if e := s.service.Announce(
		identifier,
		r.Body.Callsign,
		r.Body.Topic,
		files,
	); e != nil {
		if errors.Is(e, goclauded.ErrorCallsignNotFound) {
			return server.PostAnnounce500JSONResponse(
				server.ErrorResponse{Error: e.Error()},
			), nil
		}

		return server.PostAnnounce500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostAnnounce200Response{}, nil
}
