package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) RemoveDeviceLabel(
	_ context.Context,
	r server.RemoveDeviceLabelRequestObject,
) (server.RemoveDeviceLabelResponseObject, error) {
	if e := s.service.RemoveDeviceLabel(r.Name, r.Key); e != nil {
		return server.RemoveDeviceLabel500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.RemoveDeviceLabel204Response{}, nil
}
