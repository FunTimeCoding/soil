package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) AllocateStat(
	_ context.Context,
	r server.AllocateStatRequestObject,
) (server.AllocateStatResponseObject, error) {
	result, e := s.service.Allocate(string(r.Stat))

	if validation.Is(e) {
		return server.AllocateStat400JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.AllocateStat403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.AllocateStat503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.AllocateStat500JSONResponse(*s.answer(e)), nil
	}

	return server.AllocateStat200JSONResponse(*convert.Stats(result)), nil
}
