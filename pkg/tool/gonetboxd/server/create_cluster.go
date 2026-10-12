package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateCluster(
	_ context.Context,
	r server.CreateClusterRequestObject,
) (server.CreateClusterResponseObject, error) {
	c, e := s.service.CreateCluster(r.Body.Name, r.Body.Type, r.Body.Site)

	if e != nil {
		return server.CreateCluster500JSONResponse(*s.captureDetail(e)), nil
	}

	result := convert.Cluster(c)
	result.Site = &r.Body.Site

	return server.CreateCluster201JSONResponse(*result), nil
}
