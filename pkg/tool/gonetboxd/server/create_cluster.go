package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateCluster(
	_ context.Context,
	r server.CreateClusterRequestObject,
) (server.CreateClusterResponseObject, error) {
	c, e := s.service.CreateCluster(r.Body.Name, r.Body.Type, r.Body.Site)

	if validation.Is(e) {
		return server.CreateCluster400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.CreateCluster403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreateCluster500JSONResponse(*s.captureDetail(e)), nil
	}

	result := convert.Cluster(c)
	result.Site = &r.Body.Site

	return server.CreateCluster201JSONResponse(*result), nil
}
