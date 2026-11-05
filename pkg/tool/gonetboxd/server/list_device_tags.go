package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListDeviceTags(
	x context.Context,
	r server.ListDeviceTagsRequestObject,
) (server.ListDeviceTagsResponseObject, error) {
	tags, e := s.service.DeviceTagNames(r.Name)

	if not_found.Is(e) {
		return server.ListDeviceTags404JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.ListDeviceTags503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListDeviceTags500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.ListDeviceTags200JSONResponse(tags), nil
}
