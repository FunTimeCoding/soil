package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
)

func (s *Server) PostSessionPulse(
	x context.Context,
	r server.PostSessionPulseRequestObject,
) (server.PostSessionPulseResponseObject, error) {
	delivered, e := s.message.SendPulse(
		r.Identifier,
		"",
		r.Body.Body,
		toDelivery(r.Body.Delivery),
	)

	if e != nil {
		if refusal.Is(e) {
			return server.PostSessionPulse400JSONResponse(
				server.Error{Error: e.Error()},
			), nil
		}

		return server.PostSessionPulse500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	return server.PostSessionPulse200JSONResponse{Immediate: delivered}, nil
}
