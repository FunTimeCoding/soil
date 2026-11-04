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

func (s *Server) ShutdownMachine(
	x context.Context,
	r server.ShutdownMachineRequestObject,
) (server.ShutdownMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.ShutdownMachine404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.ShutdownMachine400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ShutdownMachine500JSONResponse(*s.answer(x, e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	taskIdentifier, e := s.service.ShutdownMachine(
		instance,
		int(r.Identifier),
		node,
	)

	switch {
	case validation.Is(e):
		return server.ShutdownMachine400JSONResponse(*clientError(e)), nil
	case not_found.Is(e):
		return server.ShutdownMachine404JSONResponse(*clientError(e)), nil
	case forbidden.Is(e):
		return server.ShutdownMachine403JSONResponse(*clientError(e)), nil
	case conflict.Is(e):
		return server.ShutdownMachine409JSONResponse(*clientError(e)), nil
	case unavailable.Is(e):
		return server.ShutdownMachine503JSONResponse(*clientError(e)), nil
	case e != nil:
		return server.ShutdownMachine500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ShutdownMachine200JSONResponse{TaskId: taskIdentifier}, nil
}
