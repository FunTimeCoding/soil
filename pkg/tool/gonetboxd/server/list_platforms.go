package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListPlatforms(
	_ context.Context,
	_ server.ListPlatformsRequestObject,
) (server.ListPlatformsResponseObject, error) {
	platforms, e := s.client.Platforms()

	if e != nil {
		return server.ListPlatforms500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListPlatforms200JSONResponse(convert.Platforms(platforms)), nil
}
