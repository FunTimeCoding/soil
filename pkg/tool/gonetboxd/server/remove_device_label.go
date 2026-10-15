package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) RemoveDeviceLabel(
	_ context.Context,
	r server.RemoveDeviceLabelRequestObject,
) (server.RemoveDeviceLabelResponseObject, error) {
	e := s.service.RemoveDeviceLabel(r.Name, r.Key)

	if not_found.Is(e) {
		return server.RemoveDeviceLabel404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.RemoveDeviceLabel500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.RemoveDeviceLabel204Response{}, nil
}
