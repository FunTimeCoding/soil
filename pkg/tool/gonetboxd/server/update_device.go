package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) UpdateDevice(
	_ context.Context,
	r server.UpdateDeviceRequestObject,
) (server.UpdateDeviceResponseObject, error) {
	if r.Body.PrimaryAddress != nil {
		if _, e := s.client.SetDevicePrimaryAddress(
			r.Name,
			*r.Body.PrimaryAddress,
		); e != nil {
			return server.UpdateDevice500JSONResponse(*s.captureDetail(e)), nil
		}
	}

	name := r.Name

	if r.Body.Name != nil {
		result, e := s.client.RenameDevice(r.Name, *r.Body.Name)

		if e != nil {
			return server.UpdateDevice500JSONResponse(*s.captureDetail(e)), nil
		}

		name = result.Name
	}

	d, e := s.client.DeviceByName(name)

	if e != nil {
		return server.UpdateDevice500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.UpdateDevice200JSONResponse(*convert.Device(d)), nil
}
