package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListDeviceLabels(
	x context.Context,
	r server.ListDeviceLabelsRequestObject,
) (server.ListDeviceLabelsResponseObject, error) {
	labels, e := s.service.DeviceLabels(r.Name)

	if not_found.Is(e) {
		return server.ListDeviceLabels404JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.ListDeviceLabels503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListDeviceLabels500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.ListDeviceLabels200JSONResponse(convert.Labels(labels)), nil
}
