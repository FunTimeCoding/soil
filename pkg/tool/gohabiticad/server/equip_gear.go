package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) EquipGear(
	_ context.Context,
	r server.EquipGearRequestObject,
) (server.EquipGearResponseObject, error) {
	result, e := s.service.Equip(r.Key)

	if not_found.Is(e) {
		return server.EquipGear404JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.EquipGear403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.EquipGear503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.EquipGear500JSONResponse(*s.answer(e)), nil
	}

	return server.EquipGear200JSONResponse(*convert.Gear(result)), nil
}
