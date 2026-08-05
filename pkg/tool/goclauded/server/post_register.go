package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostRegister(
	_ context.Context,
	r server.PostRegisterRequestObject,
) (server.PostRegisterResponseObject, error) {
	result, e := s.service.Register(r.Body.Session)

	if e != nil {
		return server.PostRegister500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostRegister200JSONResponse{
		Callsign: result.Callsign,
	}, nil
}
