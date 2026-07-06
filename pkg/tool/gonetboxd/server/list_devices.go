package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListDevices(
	_ context.Context,
	r server.ListDevicesRequestObject,
) (server.ListDevicesResponseObject, error) {
	if r.Params.Query != nil && *r.Params.Query != "" {
		devices, e := s.client.DevicesByMatch(*r.Params.Query)

		if e != nil {
			return server.ListDevices500JSONResponse(*s.captureDetail(e)), nil
		}

		return server.ListDevices200JSONResponse(convert.Devices(devices)), nil
	}

	devices, e := s.client.Devices()

	if e != nil {
		return server.ListDevices500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListDevices200JSONResponse(convert.Devices(devices)), nil
}
