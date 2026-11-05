package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateCluster(
	x context.Context,
	r server.CreateClusterRequestObject,
) (server.CreateClusterResponseObject, error) {
	c, e := s.service.CreateCluster(r.Body.Name, r.Body.Type, r.Body.Site)

	if validation.Is(e) {
		return server.CreateCluster400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreateCluster503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateCluster500JSONResponse(*s.captureDetail(x, e)), nil
	}

	result := convert.Cluster(c)
	result.Site = &r.Body.Site

	return server.CreateCluster201JSONResponse(*result), nil
}
