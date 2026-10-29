package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListWirelessNetworks(
	_ context.Context,
	_ server.ListWirelessNetworksRequestObject,
) (server.ListWirelessNetworksResponseObject, error) {
	networks, e := s.service.WirelessNetworks()

	if forbidden.Is(e) {
		return server.ListWirelessNetworks403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListWirelessNetworks500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListWirelessNetworks200JSONResponse(
			convert.WirelessNetworks(networks),
		),
		nil
}
