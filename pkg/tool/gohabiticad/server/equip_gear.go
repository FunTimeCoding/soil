package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) EquipGear(
	_ context.Context,
	r server.EquipGearRequestObject,
) (server.EquipGearResponseObject, error) {
	e := s.habitica.Equip(r.Key)

	if not_found.Is(e) {
		return server.EquipGear404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.EquipGear500JSONResponse(*s.answer(e)), nil
	}

	result, e := s.habitica.UserGear()

	if e != nil {
		return server.EquipGear500JSONResponse(*s.answer(e)), nil
	}

	return server.EquipGear200JSONResponse(*convert.Gear(result)), nil
}
