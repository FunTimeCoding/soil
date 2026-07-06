package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateClusterType(
	_ context.Context,
	r server.CreateClusterTypeRequestObject,
) (server.CreateClusterTypeResponseObject, error) {
	t, e := s.client.CreateClusterType(r.Body.Name)

	if e != nil {
		return server.CreateClusterType500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateClusterType201JSONResponse{
		Identifier: t.Identifier,
		Name:       t.Name,
	}, nil
}
