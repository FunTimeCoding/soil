package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
)

func (s *Server) PostSend(
	x context.Context,
	r server.PostSendRequestObject,
) (server.PostSendResponseObject, error) {
	var to []string

	if r.Body.To != nil {
		to = *r.Body.To
	}

	result, e := s.message.Send(
		r.Body.Callsign,
		to,
		r.Body.Body,
		toDelivery(r.Body.Delivery),
	)

	if e != nil {
		if refusal.Is(e) {
			return server.PostSend400JSONResponse(
				server.Error{Error: e.Error()},
			), nil
		}

		return server.PostSend500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	return server.PostSend200JSONResponse{Immediate: result.Immediate()}, nil
}
