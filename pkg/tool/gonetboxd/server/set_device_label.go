package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) SetDeviceLabel(
	_ context.Context,
	r server.SetDeviceLabelRequestObject,
) (server.SetDeviceLabelResponseObject, error) {
	d, e := s.client.DeviceByName(r.Name)

	if e != nil {
		return server.SetDeviceLabel500JSONResponse(*s.captureDetail(e)), nil
	}

	result, e := s.store.SetLabel(
		constant.DeviceAddress,
		d.Identifier,
		r.Key,
		r.Body.Value,
	)

	if e != nil {
		return server.SetDeviceLabel500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.SetDeviceLabel200JSONResponse(*convert.Label(result)), nil
}
