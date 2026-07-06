package server

import (
	"context"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostRegister(
	_ context.Context,
	r server.PostRegisterRequestObject,
) (server.PostRegisterResponseObject, error) {
	result, e := s.service.Register(r.Body.Session)

	if e != nil {
		return server.PostRegister500JSONResponse(
			*s.captureFail(e, library.UnexpectedError),
		), nil
	}

	s.logger.Structured(
		constant.Register,
		"claude_session_identifier",
		r.Body.Session,
		constant.SessionName,
		result.Callsign,
		"new",
		result.New,
	)

	return server.PostRegister200JSONResponse{
		Callsign: result.Callsign,
	}, nil
}
