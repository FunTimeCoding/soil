package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) UpdateMachine(
	x context.Context,
	r server.UpdateMachineRequestObject,
) (server.UpdateMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.UpdateMachine404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.UpdateMachine400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.UpdateMachine500JSONResponse(*s.answer(x, e)), nil
	}

	a := convertUpdateMachine(r)
	e = s.service.UpdateMachine(instance, a)

	switch {
	case validation.Is(e):
		return server.UpdateMachine400JSONResponse(*clientError(e)), nil
	case not_found.Is(e):
		return server.UpdateMachine404JSONResponse(*clientError(e)), nil
	case forbidden.Is(e):
		return server.UpdateMachine403JSONResponse(*clientError(e)), nil
	case conflict.Is(e):
		return server.UpdateMachine409JSONResponse(*clientError(e)), nil
	case unavailable.Is(e):
		return server.UpdateMachine503JSONResponse(*clientError(e)), nil
	case e != nil:
		return server.UpdateMachine500JSONResponse(*s.answer(x, e)), nil
	}

	status := "updated"

	return server.UpdateMachine200JSONResponse{Status: status}, nil
}
