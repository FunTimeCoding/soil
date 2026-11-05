package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListTenants(
	x context.Context,
	_ server.ListTenantsRequestObject,
) (server.ListTenantsResponseObject, error) {
	tenants, e := s.service.Tenants()

	if unavailable.Is(e) {
		return server.ListTenants503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListTenants500JSONResponse(*s.captureDetail(x, e)), nil
	}

	result := make([]*server.Tenant, 0, len(tenants))

	for _, t := range tenants {
		result = append(
			result,
			&server.Tenant{Identifier: t.Identifier, Name: t.Name},
		)
	}

	return server.ListTenants200JSONResponse(result), nil
}
