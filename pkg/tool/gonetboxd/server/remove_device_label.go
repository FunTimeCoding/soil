package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) RemoveDeviceLabel(
	x context.Context,
	r server.RemoveDeviceLabelRequestObject,
) (server.RemoveDeviceLabelResponseObject, error) {
	e := s.service.RemoveDeviceLabel(r.Name, r.Key)

	if not_found.Is(e) {
		return server.RemoveDeviceLabel404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.RemoveDeviceLabel400JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.RemoveDeviceLabel409JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.RemoveDeviceLabel503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.RemoveDeviceLabel500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.RemoveDeviceLabel204Response{}, nil
}
