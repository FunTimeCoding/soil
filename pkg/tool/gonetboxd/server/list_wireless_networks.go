package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListWirelessNetworks(
	x context.Context,
	_ server.ListWirelessNetworksRequestObject,
) (server.ListWirelessNetworksResponseObject, error) {
	networks, e := s.service.WirelessNetworks()

	if unavailable.Is(e) {
		return server.ListWirelessNetworks503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListWirelessNetworks500JSONResponse(
			*s.captureDetail(x, e),
		), nil
	}

	return server.ListWirelessNetworks200JSONResponse(
			convert.WirelessNetworks(networks),
		),
		nil
}
