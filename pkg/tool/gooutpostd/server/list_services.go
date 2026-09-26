package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/server"
)

func (s *Server) ListServices(
	_ context.Context,
	r server.ListServicesRequestObject,
) (server.ListServicesResponseObject, error) {
	origin := ""

	if r.Params.Origin != nil {
		origin = *r.Params.Origin
	}

	return server.ListServices200JSONResponse(
		convert.Services(s.service.Services(), origin),
	), nil
}
