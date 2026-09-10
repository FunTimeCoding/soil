package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostSessionEnd(
	_ context.Context,
	r server.PostSessionEndRequestObject,
) (server.PostSessionEndResponseObject, error) {
	reason := ""

	if r.Body.Reason != nil {
		reason = *r.Body.Reason
	}

	if e := s.service.CloseSession(r.Body.Session, reason); e != nil {
		return server.PostSessionEnd500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostSessionEnd200Response{}, nil
}
