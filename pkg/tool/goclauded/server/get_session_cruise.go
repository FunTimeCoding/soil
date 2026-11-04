package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetSessionCruise(
	x context.Context,
	r server.GetSessionCruiseRequestObject,
) (server.GetSessionCruiseResponseObject, error) {
	c, e := s.message.CruiseSetting(r.Identifier)

	if e != nil {
		return server.GetSessionCruise500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	return server.GetSessionCruise200JSONResponse(toCruiseResponse(c)), nil
}
