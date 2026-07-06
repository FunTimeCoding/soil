package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) GetDevice(
	_ context.Context,
	r server.GetDeviceRequestObject,
) (server.GetDeviceResponseObject, error) {
	d, e := s.client.DeviceByName(r.Name)

	if e != nil {
		return server.GetDevice500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.GetDevice200JSONResponse(*convert.Device(d)), nil
}
