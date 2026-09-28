package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/server"
)

func (s *Server) ListSightings(
	_ context.Context,
	r server.ListSightingsRequestObject,
) (server.ListSightingsResponseObject, error) {
	var place string

	if r.Params.Place != nil {
		place = *r.Params.Place
	}

	v, e := s.store.MatchingSightings(
		place,
		r.Params.Unclaimed != nil && *r.Params.Unclaimed,
	)

	if e != nil {
		return server.ListSightings500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.ListSightings200JSONResponse(convert.SightingSlice(v)), nil
}
