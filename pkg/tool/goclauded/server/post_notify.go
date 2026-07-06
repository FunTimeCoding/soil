package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostNotify(
	_ context.Context,
	r server.PostNotifyRequestObject,
) (server.PostNotifyResponseObject, error) {
	if e := s.service.SendNotification(
		r.Body.Callsign,
		r.Body.Source,
		r.Body.Body,
	); e != nil {
		return server.PostNotify500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostNotify200Response{}, nil
}
