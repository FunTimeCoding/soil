package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateWirelessNetwork(
	_ context.Context,
	r server.CreateWirelessNetworkRequestObject,
) (server.CreateWirelessNetworkResponseObject, error) {
	result, e := s.service.CreateWirelessNetwork(r.Body.Name)

	if validation.Is(e) {
		return server.CreateWirelessNetwork400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.CreateWirelessNetwork403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreateWirelessNetwork500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateWirelessNetwork201JSONResponse(
			*convert.WirelessNetwork(result),
		),
		nil
}
