package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
	"time"
)

func (s *Server) PostSessionCruise(
	_ context.Context,
	r server.PostSessionCruiseRequestObject,
) (server.PostSessionCruiseResponseObject, error) {
	var pace time.Duration

	if r.Body.PaceMinutes != nil {
		pace = time.Duration(*r.Body.PaceMinutes) * time.Minute
	}

	if e := s.service.SetCruise(
		r.Identifier,
		constant.Cruise(r.Body.Mode),
		pace,
	); e != nil {
		if refusal.Is(e) {
			return server.PostSessionCruise400JSONResponse(
				server.Error{Error: e.Error()},
			), nil
		}

		return server.PostSessionCruise500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	c, f := s.service.CruiseSetting(r.Identifier)

	if f != nil {
		return server.PostSessionCruise500JSONResponse(
			*s.captureFail(f, soil.UnexpectedError),
		), nil
	}

	return server.PostSessionCruise200JSONResponse(toCruiseResponse(c)), nil
}
