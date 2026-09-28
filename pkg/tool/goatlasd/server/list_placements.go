package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/server"
)

func (s *Server) ListPlacements(
	_ context.Context,
	r server.ListPlacementsRequestObject,
) (server.ListPlacementsResponseObject, error) {
	var place string
	var source string
	var packageName string

	if r.Params.Place != nil {
		place = *r.Params.Place
	}

	if r.Params.Source != nil {
		source = *r.Params.Source
	}

	if r.Params.Package != nil {
		packageName = *r.Params.Package
	}

	v, e := s.store.MatchingPlacements(place, source, packageName)

	if e != nil {
		return server.ListPlacements500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.ListPlacements200JSONResponse(convert.PlacementSlice(v)), nil
}
