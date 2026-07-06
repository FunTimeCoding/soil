package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListManufacturers(
	_ context.Context,
	_ server.ListManufacturersRequestObject,
) (server.ListManufacturersResponseObject, error) {
	manufacturers, e := s.client.Manufacturers()

	if e != nil {
		return server.ListManufacturers500JSONResponse(*s.captureDetail(e)), nil
	}

	result := make([]*server.Manufacturer, 0, len(manufacturers))

	for _, m := range manufacturers {
		result = append(
			result,
			&server.Manufacturer{
				Identifier: m.Identifier, Name: m.Name,
			},
		)
	}

	return server.ListManufacturers200JSONResponse(result), nil
}
