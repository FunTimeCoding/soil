package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateLocation(
	x context.Context,
	r server.CreateLocationRequestObject,
) (server.CreateLocationResponseObject, error) {
	result, e := s.service.CreateLocation(r.Body.Name, r.Body.Site)

	if validation.Is(e) {
		return server.CreateLocation400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreateLocation503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateLocation500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.CreateLocation201JSONResponse(*convert.Location(result)), nil
}
