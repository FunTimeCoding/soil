package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) RemoveDeviceTag(
	_ context.Context,
	r server.RemoveDeviceTagRequestObject,
) (server.RemoveDeviceTagResponseObject, error) {
	d, e := s.client.RemoveTag(r.Name, r.Tag)

	if not_found.Is(e) {
		return server.RemoveDeviceTag404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.RemoveDeviceTag500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.RemoveDeviceTag200JSONResponse(*convert.Device(d)), nil
}
