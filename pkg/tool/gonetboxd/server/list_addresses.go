package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListAddresses(
	x context.Context,
	r server.ListAddressesRequestObject,
) (server.ListAddressesResponseObject, error) {
	addresses, e := s.service.DeviceAddresses(r.Name)

	if not_found.Is(e) {
		return server.ListAddresses404JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.ListAddresses503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListAddresses500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.ListAddresses200JSONResponse(convert.Addresses(addresses)), nil
}
