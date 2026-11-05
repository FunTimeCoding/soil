package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateClusterType(
	x context.Context,
	r server.CreateClusterTypeRequestObject,
) (server.CreateClusterTypeResponseObject, error) {
	t, e := s.service.CreateClusterType(r.Body.Name)

	if validation.Is(e) {
		return server.CreateClusterType400JSONResponse(*clientError(e)),
			nil
	}

	if unavailable.Is(e) {
		return server.CreateClusterType503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateClusterType500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.CreateClusterType201JSONResponse{
		Identifier: t.Identifier,
		Name:       t.Name,
	}, nil
}
