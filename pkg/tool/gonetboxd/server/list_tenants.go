package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListTenants(
	_ context.Context,
	_ server.ListTenantsRequestObject,
) (server.ListTenantsResponseObject, error) {
	tenants, e := s.service.Tenants()

	if forbidden.Is(e) {
		return server.ListTenants403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListTenants500JSONResponse(*s.captureDetail(e)), nil
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
