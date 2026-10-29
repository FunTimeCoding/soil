package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListPlatforms(
	_ context.Context,
	_ server.ListPlatformsRequestObject,
) (server.ListPlatformsResponseObject, error) {
	platforms, e := s.service.Platforms()

	if forbidden.Is(e) {
		return server.ListPlatforms403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListPlatforms500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListPlatforms200JSONResponse(convert.Platforms(platforms)), nil
}
