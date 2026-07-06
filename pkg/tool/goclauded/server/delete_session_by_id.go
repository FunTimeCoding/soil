package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) DeleteSessionById(
	_ context.Context,
	r server.DeleteSessionByIdRequestObject,
) (server.DeleteSessionByIdResponseObject, error) {
	if e := s.service.DeleteSession(r.Identifier); e != nil {
		return server.DeleteSessionById500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.DeleteSessionById200Response{}, nil
}
