package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListDeviceTags(
	_ context.Context,
	r server.ListDeviceTagsRequestObject,
) (server.ListDeviceTagsResponseObject, error) {
	tags, e := s.service.DeviceTagNames(r.Name)

	if not_found.Is(e) {
		return server.ListDeviceTags404JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.ListDeviceTags403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListDeviceTags500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListDeviceTags200JSONResponse(tags), nil
}
