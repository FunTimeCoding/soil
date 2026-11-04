package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/server"
)

func (s *Server) ListServices(
	x context.Context,
	r server.ListServicesRequestObject,
) (server.ListServicesResponseObject, error) {
	var origin string

	if r.Params.Origin != nil {
		origin = *r.Params.Origin
	}

	v, e := s.service.Services(origin)

	if e != nil {
		return server.ListServices500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ListServices200JSONResponse(convert.Services(v)), nil
}
