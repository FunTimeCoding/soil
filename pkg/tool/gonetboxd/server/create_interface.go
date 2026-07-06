package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"github.com/netbox-community/go-netbox/v4"
)

func (s *Server) CreateInterface(
	_ context.Context,
	r server.CreateInterfaceRequestObject,
) (server.CreateInterfaceResponseObject, error) {
	d, e := s.client.DeviceByName(r.Name)

	if e != nil {
		return server.CreateInterface500JSONResponse(*s.captureDetail(e)), nil
	}

	i, f := s.client.CreateInterface(
		d,
		r.Body.Name,
		netbox.InterfaceTypeValue(r.Body.Type),
	)

	if f != nil {
		return server.CreateInterface500JSONResponse(*s.captureDetail(f)), nil
	}

	result := server.Interface{
		Identifier: i.Identifier,
		Name:       i.Name,
	}

	if i.Type != "" {
		t := string(i.Type)
		result.Type = &t
	}

	return server.CreateInterface201JSONResponse(result), nil
}
