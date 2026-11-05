package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListManufacturers(
	x context.Context,
	_ server.ListManufacturersRequestObject,
) (server.ListManufacturersResponseObject, error) {
	manufacturers, e := s.service.Manufacturers()

	if unavailable.Is(e) {
		return server.ListManufacturers503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListManufacturers500JSONResponse(*s.captureDetail(x, e)), nil
	}

	result := make([]*server.Manufacturer, 0, len(manufacturers))

	for _, m := range manufacturers {
		result = append(
			result,
			&server.Manufacturer{Identifier: m.Identifier, Name: m.Name},
		)
	}

	return server.ListManufacturers200JSONResponse(result), nil
}
