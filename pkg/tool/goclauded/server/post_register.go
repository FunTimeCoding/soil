package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostRegister(
	x context.Context,
	r server.PostRegisterRequestObject,
) (server.PostRegisterResponseObject, error) {
	result, e := s.roster.Register(r.Body.Session)

	if e != nil {
		return server.PostRegister500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	return server.PostRegister200JSONResponse{Callsign: result.Callsign}, nil
}
