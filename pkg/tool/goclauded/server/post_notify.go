package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
)

func (s *Server) PostNotify(
	_ context.Context,
	r server.PostNotifyRequestObject,
) (server.PostNotifyResponseObject, error) {
	delivered, e := s.message.SendNotification(
		r.Body.Callsign,
		r.Body.Source,
		r.Body.Body,
		toDelivery(r.Body.Delivery),
	)

	if e != nil {
		if refusal.Is(e) {
			return server.PostNotify400JSONResponse(
				server.Error{Error: e.Error()},
			), nil
		}

		if not_found.Is(e) {
			return server.PostNotify404JSONResponse(
				server.Error{Error: e.Error()},
			), nil
		}

		return server.PostNotify500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	return server.PostNotify200JSONResponse{Immediate: delivered}, nil
}
