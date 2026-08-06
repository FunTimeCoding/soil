package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateWirelessNetwork(
	_ context.Context,
	r server.CreateWirelessNetworkRequestObject,
) (server.CreateWirelessNetworkResponseObject, error) {
	result, e := s.client.CreateWirelessNetwork(r.Body.Name)

	if e != nil {
		return server.CreateWirelessNetwork500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.CreateWirelessNetwork201JSONResponse(
		*convert.WirelessNetwork(result),
	), nil
}
