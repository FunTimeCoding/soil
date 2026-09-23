package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostTurnEnd(
	_ context.Context,
	r server.PostTurnEndRequestObject,
) (server.PostTurnEndResponseObject, error) {
	if e := s.service.StampTurnEnd(r.Body.Session); e != nil {
		return server.PostTurnEnd500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostTurnEnd200Response{}, nil
}
