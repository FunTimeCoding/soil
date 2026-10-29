package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListDevices(
	_ context.Context,
	r server.ListDevicesRequestObject,
) (server.ListDevicesResponseObject, error) {
	query := ""

	if r.Params.Query != nil {
		query = *r.Params.Query
	}

	devices, grouped, e := s.service.Devices(query)

	if forbidden.Is(e) {
		return server.ListDevices403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListDevices500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListDevices200JSONResponse(
			convert.LabeledDevices(devices, grouped),
		),
		nil
}
