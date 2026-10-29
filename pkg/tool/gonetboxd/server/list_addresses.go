package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListAddresses(
	_ context.Context,
	r server.ListAddressesRequestObject,
) (server.ListAddressesResponseObject, error) {
	addresses, e := s.service.DeviceAddresses(r.Name)

	if not_found.Is(e) {
		return server.ListAddresses404JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.ListAddresses403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListAddresses500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListAddresses200JSONResponse(convert.Addresses(addresses)), nil
}
