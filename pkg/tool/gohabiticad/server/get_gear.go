package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) GetGear(
	_ context.Context,
	_ server.GetGearRequestObject,
) (server.GetGearResponseObject, error) {
	result, e := s.service.UserGear()

	if forbidden.Is(e) {
		return server.GetGear403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetGear503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetGear500JSONResponse(*s.answer(e)), nil
	}

	return server.GetGear200JSONResponse(*convert.Gear(result)), nil
}
