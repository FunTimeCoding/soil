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
	deviceA, e := s.client.DeviceByName(r.Body.DeviceA)

	if e != nil {
		return server.CreateCable500JSONResponse(*s.captureDetail(e)), nil
	}

	sideA, f := s.client.DeviceInterfaceByName(deviceA, r.Body.InterfaceA)

	if f != nil {
		return server.CreateCable500JSONResponse(*s.captureDetail(f)), nil
	}

	deviceB, g := s.client.DeviceByName(r.Body.DeviceB)

	if g != nil {
		return server.CreateCable500JSONResponse(*s.captureDetail(g)), nil
	}

	sideB, h := s.client.DeviceInterfaceByName(deviceB, r.Body.InterfaceB)

	if h != nil {
		return server.CreateCable500JSONResponse(*s.captureDetail(h)), nil
	}

	result, createFail := s.client.CreateCable(sideA, sideB)

	if createFail != nil {
		return server.CreateCable500JSONResponse(
			*s.captureDetail(createFail),
		), nil
	}

	return server.CreateCable201JSONResponse(*convert.Cable(result)), nil
}
