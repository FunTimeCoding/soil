package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/server"
)

func (s *Server) ListPlaces(
	_ context.Context,
	_ server.ListPlacesRequestObject,
) (server.ListPlacesResponseObject, error) {
	v, e := s.store.Places()

	if e != nil {
		return server.ListPlaces500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.ListPlaces200JSONResponse(convert.PlaceSlice(v)), nil
}
