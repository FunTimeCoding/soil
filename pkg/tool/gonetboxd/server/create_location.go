package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateLocation(
	_ context.Context,
	r server.CreateLocationRequestObject,
) (server.CreateLocationResponseObject, error) {
	result, e := s.client.CreateLocation(r.Body.Name, r.Body.Site)

	if e != nil {
		return server.CreateLocation500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateLocation201JSONResponse(
		*convert.Location(result),
	), nil
}
