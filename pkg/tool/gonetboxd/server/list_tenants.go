package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
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

	return server.ListTenants200JSONResponse(convert.Tenants(tenants)), nil
}
