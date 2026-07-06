package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostSessionPulse(
	_ context.Context,
	r server.PostSessionPulseRequestObject,
) (server.PostSessionPulseResponseObject, error) {
	if e := s.service.SendPulse(r.Identifier, "", r.Body.Body); e != nil {
		return server.PostSessionPulse500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostSessionPulse200Response{}, nil
}
