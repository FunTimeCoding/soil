package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListPowerOutlets(
	x context.Context,
	r server.ListPowerOutletsRequestObject,
) (server.ListPowerOutletsResponseObject, error) {
	outlets, e := s.service.DevicePowerOutlets(r.Name)

	if not_found.Is(e) {
		return server.ListPowerOutlets404JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.ListPowerOutlets503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListPowerOutlets500JSONResponse(
			*s.captureDetail(x, e),
		), nil
	}

	return server.ListPowerOutlets200JSONResponse(
		convert.PowerOutlets(outlets),
	), nil
}
