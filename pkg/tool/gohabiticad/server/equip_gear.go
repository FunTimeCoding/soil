package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) EquipGear(
	_ context.Context,
	r server.EquipGearRequestObject,
) (server.EquipGearResponseObject, error) {
	if e := s.habitica.Equip(r.Key); e != nil {
		return server.EquipGear500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	result, e := s.habitica.UserGear()

	if e != nil {
		return server.EquipGear500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.EquipGear200JSONResponse(
		*convert.Gear(result),
	), nil
}
