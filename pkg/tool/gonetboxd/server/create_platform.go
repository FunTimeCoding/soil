package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreatePlatform(
	x context.Context,
	r server.CreatePlatformRequestObject,
) (server.CreatePlatformResponseObject, error) {
	result, e := s.service.CreatePlatform(r.Body.Name)

	if validation.Is(e) {
		return server.CreatePlatform400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreatePlatform503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreatePlatform500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.CreatePlatform201JSONResponse(*convert.Platform(result)), nil
}
