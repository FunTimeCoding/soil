package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListDeviceLabels(
	_ context.Context,
	r server.ListDeviceLabelsRequestObject,
) (server.ListDeviceLabelsResponseObject, error) {
	d, e := s.client.DeviceByName(r.Name)

	if e != nil {
		return server.ListDeviceLabels500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	labels, e := s.store.Labels(constant.DeviceAddress, d.Identifier)

	if e != nil {
		return server.ListDeviceLabels500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.ListDeviceLabels200JSONResponse(
		convert.Labels(labels),
	), nil
}
