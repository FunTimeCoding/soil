package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) CreateMachine(
	_ context.Context,
	r server.CreateMachineRequestObject,
) (server.CreateMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.CreateMachine404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.CreateMachine400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreateMachine500JSONResponse(*s.captureFail(e, e.Error())), nil
	}

	m := convertCreateMachine(r.Body)
	identifier, e := s.service.CreateMachine(instance, m)

	if e != nil {
		if validation.Is(e) {
			return server.CreateMachine400JSONResponse(*clientError(e)), nil
		}

		return server.CreateMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	status := "created"

	return server.CreateMachine200JSONResponse{
		Identifier: int64(identifier),
		Status:     &status,
	}, nil
}
