package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateTenant(
	_ context.Context,
	r server.CreateTenantRequestObject,
) (server.CreateTenantResponseObject, error) {
	t, e := s.service.CreateTenant(r.Body.Name)

	if validation.Is(e) {
		return server.CreateTenant400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.CreateTenant403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreateTenant500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateTenant201JSONResponse{
		Identifier: t.Identifier,
		Name:       t.Name,
	}, nil
}
