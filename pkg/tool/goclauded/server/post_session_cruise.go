package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
)

func (s *Server) PostSessionCruise(
	x context.Context,
	r server.PostSessionCruiseRequestObject,
) (server.PostSessionCruiseResponseObject, error) {
	if e := s.message.SetCruise(r.Identifier, r.Body.Grace); e != nil {
		if refusal.Is(e) {
			return server.PostSessionCruise400JSONResponse(
				server.Error{Error: e.Error()},
			), nil
		}

		return server.PostSessionCruise500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	c, f := s.message.CruiseSetting(r.Identifier)

	if f != nil {
		return server.PostSessionCruise500JSONResponse(
			*s.captureFail(x, f, soil.UnexpectedError),
		), nil
	}

	return server.PostSessionCruise200JSONResponse(
		toCruiseResponse(c, s.message.Grace(c)),
	), nil
}
