package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateTenant(
	_ context.Context,
	r server.CreateTenantRequestObject,
) (server.CreateTenantResponseObject, error) {
	t, e := s.client.CreateTenant(r.Body.Name)

	if e != nil {
		return server.CreateTenant500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateTenant201JSONResponse{
		Identifier: t.Identifier, Name: t.Name,
	}, nil
}
