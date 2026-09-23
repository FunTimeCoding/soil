package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetChannelCallsign(
	x context.Context,
	r server.GetChannelCallsignRequestObject,
) (server.GetChannelCallsignResponseObject, error) {
	callsign, e := s.service.AwaitCallsign(x, r.Params.Session, r.Params.Since)

	if e != nil {
		return server.GetChannelCallsign500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	if callsign == "" {
		return server.GetChannelCallsign204Response{}, nil
	}

	return server.GetChannelCallsign200JSONResponse(
		server.ChannelCallsignResponse{Callsign: callsign},
	), nil
}
