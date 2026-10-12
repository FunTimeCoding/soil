package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateVirtualAddress(
	_ context.Context,
	r server.CreateVirtualAddressRequestObject,
) (server.CreateVirtualAddressResponseObject, error) {
	status := ""

	if r.Body.Status != nil {
		status = *r.Body.Status
	}

	a, e := s.service.CreateVirtualAddress(
		r.Name,
		r.Body.Interface,
		r.Body.Address,
		status,
	)

	if e != nil {
		return server.CreateVirtualAddress500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateVirtualAddress201JSONResponse{
		Identifier: a.Identifier,
		Address:    a.Address.String(),
	}, nil
}
