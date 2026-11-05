package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListPlatforms(
	x context.Context,
	_ server.ListPlatformsRequestObject,
) (server.ListPlatformsResponseObject, error) {
	platforms, e := s.service.Platforms()

	if unavailable.Is(e) {
		return server.ListPlatforms503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListPlatforms500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.ListPlatforms200JSONResponse(convert.Platforms(platforms)), nil
}
