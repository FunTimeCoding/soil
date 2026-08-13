package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListLocations(
	_ context.Context,
	_ server.ListLocationsRequestObject,
) (server.ListLocationsResponseObject, error) {
	locations, e := s.client.Locations()

	if e != nil {
		return server.ListLocations500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListLocations200JSONResponse(convert.Locations(locations)), nil
}
