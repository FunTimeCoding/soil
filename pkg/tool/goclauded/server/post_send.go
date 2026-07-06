package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostSend(
	_ context.Context,
	r server.PostSendRequestObject,
) (server.PostSendResponseObject, error) {
	var to string

	if r.Body.To != nil {
		to = *r.Body.To
	}

	if e := s.service.Send(r.Body.Callsign, to, r.Body.Body); e != nil {
		return server.PostSend500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostSend200Response{}, nil
}
