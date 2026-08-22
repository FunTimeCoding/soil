package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) UpdateMachine(
	_ context.Context,
	r server.UpdateMachineRequestObject,
) (server.UpdateMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.UpdateMachine400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.UpdateMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	a := convertUpdateMachine(r)
	e = s.service.UpdateMachine(c, a)

	if e != nil {
		if not_found.Is(e) {
			return server.UpdateMachine404JSONResponse{Error: e.Error()}, nil
		}

		if validation.Is(e) {
			return server.UpdateMachine400JSONResponse(*clientError(e)), nil
		}

		return server.UpdateMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	status := "updated"

	return server.UpdateMachine200JSONResponse{Status: status}, nil
}
