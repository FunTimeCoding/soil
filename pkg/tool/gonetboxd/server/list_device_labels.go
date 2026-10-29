package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListDeviceLabels(
	_ context.Context,
	r server.ListDeviceLabelsRequestObject,
) (server.ListDeviceLabelsResponseObject, error) {
	labels, e := s.service.DeviceLabels(r.Name)

	if not_found.Is(e) {
		return server.ListDeviceLabels404JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.ListDeviceLabels403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListDeviceLabels500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListDeviceLabels200JSONResponse(convert.Labels(labels)), nil
}
