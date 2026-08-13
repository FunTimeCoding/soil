package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreatePlatform(
	_ context.Context,
	r server.CreatePlatformRequestObject,
) (server.CreatePlatformResponseObject, error) {
	result, e := s.client.CreatePlatform(r.Body.Name)

	if e != nil {
		return server.CreatePlatform500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreatePlatform201JSONResponse(*convert.Platform(result)), nil
}
