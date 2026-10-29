package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateClusterType(
	_ context.Context,
	r server.CreateClusterTypeRequestObject,
) (server.CreateClusterTypeResponseObject, error) {
	t, e := s.service.CreateClusterType(r.Body.Name)

	if validation.Is(e) {
		return server.CreateClusterType400JSONResponse(*clientError(e)),
			nil
	}

	if forbidden.Is(e) {
		return server.CreateClusterType403JSONResponse(*clientError(e)),
			nil
	}

	if e != nil {
		return server.CreateClusterType500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateClusterType201JSONResponse{
		Identifier: t.Identifier,
		Name:       t.Name,
	}, nil
}
