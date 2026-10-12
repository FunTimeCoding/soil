package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateCable(
	_ context.Context,
	r server.CreateCableRequestObject,
) (server.CreateCableResponseObject, error) {
	result, e := s.service.CreateCable(
		r.Body.DeviceA,
		r.Body.InterfaceA,
		r.Body.DeviceB,
		r.Body.InterfaceB,
	)

	if e != nil {
		return server.CreateCable500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateCable201JSONResponse(*convert.Cable(result)), nil
}
