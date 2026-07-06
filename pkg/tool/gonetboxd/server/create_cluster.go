package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateCluster(
	_ context.Context,
	r server.CreateClusterRequestObject,
) (server.CreateClusterResponseObject, error) {
	t, e := s.client.ClusterTypeByName(r.Body.Type)

	if e != nil {
		return server.CreateCluster500JSONResponse(*s.captureDetail(e)), nil
	}

	i, f := s.client.SiteByName(r.Body.Site)

	if f != nil {
		return server.CreateCluster500JSONResponse(*s.captureDetail(f)), nil
	}

	cl, g := s.client.CreateCluster(r.Body.Name, t, i)

	if g != nil {
		return server.CreateCluster500JSONResponse(*s.captureDetail(g)), nil
	}

	return server.CreateCluster201JSONResponse{
		Identifier: cl.Identifier,
		Name:       cl.Name,
		Type:       &t.Name,
		Site:       &i.Name,
	}, nil
}
