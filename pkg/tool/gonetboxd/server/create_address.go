package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateAddress(
	_ context.Context,
	r server.CreateAddressRequestObject,
) (server.CreateAddressResponseObject, error) {
	d, e := s.client.DeviceByName(r.Name)

	if e != nil {
		return server.CreateAddress500JSONResponse(*s.captureDetail(e)), nil
	}

	i, f := s.client.DeviceInterfaceByName(d, r.Body.Interface)

	if f != nil {
		return server.CreateAddress500JSONResponse(*s.captureDetail(f)), nil
	}

	a, g := s.client.CreateAddress(i.Identifier, r.Body.Address)

	if g != nil {
		return server.CreateAddress500JSONResponse(*s.captureDetail(g)), nil
	}

	return server.CreateAddress201JSONResponse{
		Identifier: a.Identifier,
		Address:    a.Address.String(),
	}, nil
}
