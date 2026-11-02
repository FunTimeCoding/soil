package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostTurnEnd(
	_ context.Context,
	r server.PostTurnEndRequestObject,
) (server.PostTurnEndResponseObject, error) {
	if e := s.roster.StampTurnEnd(r.Body.Session); e != nil {
		return server.PostTurnEnd500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	return server.PostTurnEnd200Response{}, nil
}
