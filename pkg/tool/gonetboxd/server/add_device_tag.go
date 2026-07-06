package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) AddDeviceTag(
	_ context.Context,
	r server.AddDeviceTagRequestObject,
) (server.AddDeviceTagResponseObject, error) {
	d, e := s.client.AddTag(r.Name, r.Tag)

	if e != nil {
		return server.AddDeviceTag500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.AddDeviceTag200JSONResponse(*convert.Device(d)), nil
}
