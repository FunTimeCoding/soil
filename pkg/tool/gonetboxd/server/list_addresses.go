package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListAddresses(
	_ context.Context,
	r server.ListAddressesRequestObject,
) (server.ListAddressesResponseObject, error) {
	addresses, e := s.client.DeviceAddresses(r.Name)

	if e != nil {
		return server.ListAddresses500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListAddresses200JSONResponse(
		convert.Addresses(addresses),
	), nil
}
