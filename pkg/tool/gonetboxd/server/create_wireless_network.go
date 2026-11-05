package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateWirelessNetwork(
	x context.Context,
	r server.CreateWirelessNetworkRequestObject,
) (server.CreateWirelessNetworkResponseObject, error) {
	result, e := s.service.CreateWirelessNetwork(r.Body.Name)

	if validation.Is(e) {
		return server.CreateWirelessNetwork400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreateWirelessNetwork503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateWirelessNetwork500JSONResponse(
			*s.captureDetail(x, e),
		), nil
	}

	return server.CreateWirelessNetwork201JSONResponse(
			*convert.WirelessNetwork(result),
		),
		nil
}
