package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetSessionCruise(
	_ context.Context,
	r server.GetSessionCruiseRequestObject,
) (server.GetSessionCruiseResponseObject, error) {
	c, e := s.service.CruiseSetting(r.Identifier)

	if e != nil {
		return server.GetSessionCruise500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.GetSessionCruise200JSONResponse(toCruiseResponse(c)), nil
}
