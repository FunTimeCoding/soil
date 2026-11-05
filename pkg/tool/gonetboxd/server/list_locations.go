package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListLocations(
	x context.Context,
	_ server.ListLocationsRequestObject,
) (server.ListLocationsResponseObject, error) {
	locations, e := s.service.Locations()

	if unavailable.Is(e) {
		return server.ListLocations503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListLocations500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.ListLocations200JSONResponse(convert.Locations(locations)), nil
}
