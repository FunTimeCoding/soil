package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) CreateMachine(
	_ context.Context,
	r server.CreateMachineRequestObject,
) (server.CreateMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.CreateMachine400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.CreateMachine500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	m := convertCreateMachine(r.Body)
	identifier, e := s.service.CreateMachine(c, m)

	if e != nil {
		if errors.Is(e, constant.ErrorCDROMCloudInitConflict) {
			return server.CreateMachine400JSONResponse(
				*clientError(e),
			), nil
		}

		return server.CreateMachine500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	status := "created"

	return server.CreateMachine200JSONResponse{
		Identifier: int64(identifier),
		Status:     &status,
	}, nil
}
