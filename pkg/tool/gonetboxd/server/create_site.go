package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateSite(
	_ context.Context,
	r server.CreateSiteRequestObject,
) (server.CreateSiteResponseObject, error) {
	i, e := s.client.CreateSite(r.Body.Name)

	if e != nil {
		return server.CreateSite500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateSite201JSONResponse{
		Identifier: i.Identifier, Name: i.Name,
	}, nil
}
