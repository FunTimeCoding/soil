package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListDeviceTags(
	_ context.Context,
	r server.ListDeviceTagsRequestObject,
) (server.ListDeviceTagsResponseObject, error) {
	tags, e := s.client.DeviceTagNames(r.Name)

	if e != nil {
		return server.ListDeviceTags500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListDeviceTags200JSONResponse(tags), nil
}
