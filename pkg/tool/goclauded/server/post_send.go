package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostSend(
	_ context.Context,
	r server.PostSendRequestObject,
) (server.PostSendResponseObject, error) {
	var to []string

	if r.Body.To != nil {
		to = *r.Body.To
	}

	d := toDelivery(r.Body.Delivery)
	held, e := s.service.Send(r.Body.Callsign, to, r.Body.Body, d)

	if e != nil {
		return server.PostSend500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	return server.PostSend200JSONResponse{
		Immediate: len(to) > 0 && d != constant.DeliveryQueued && len(held) == 0,
	}, nil
}
