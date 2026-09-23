package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostNotify(
	_ context.Context,
	r server.PostNotifyRequestObject,
) (server.PostNotifyResponseObject, error) {
	delivered, e := s.service.SendNotification(
		r.Body.Callsign,
		r.Body.Source,
		r.Body.Body,
		r.Body.Immediate != nil && *r.Body.Immediate,
	)

	if e != nil {
		if not_found.Is(e) {
			return server.PostNotify404JSONResponse(
				server.Error{Error: e.Error()},
			), nil
		}

		return server.PostNotify500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostNotify200JSONResponse{Immediate: delivered}, nil
}
