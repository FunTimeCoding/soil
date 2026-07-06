package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) GetGear(
	_ context.Context,
	_ server.GetGearRequestObject,
) (server.GetGearResponseObject, error) {
	result, e := s.habitica.UserGear()

	if e != nil {
		return server.GetGear500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.GetGear200JSONResponse(
		*convert.Gear(result),
	), nil
}
