package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) RemoveDeviceLabel(
	_ context.Context,
	r server.RemoveDeviceLabelRequestObject,
) (server.RemoveDeviceLabelResponseObject, error) {
	d, e := s.client.DeviceByName(r.Name)

	if e != nil {
		return server.RemoveDeviceLabel500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	if e := s.store.RemoveLabel(
		constant.DeviceAddress,
		d.Identifier,
		r.Key,
	); e != nil {
		return server.RemoveDeviceLabel500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.RemoveDeviceLabel204Response{}, nil
}
