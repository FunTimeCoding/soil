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

func (s *Server) DeleteMachineSnapshot(
	_ context.Context,
	r server.DeleteMachineSnapshotRequestObject,
) (server.DeleteMachineSnapshotResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.DeleteMachineSnapshot404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.DeleteMachineSnapshot400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.DeleteMachineSnapshot500JSONResponse(*s.answer(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	taskIdentifier, e := s.service.DeleteMachineSnapshot(
		instance,
		int(r.Identifier),
		node,
		r.Name,
	)

	switch {
	case validation.Is(e):
		return server.DeleteMachineSnapshot400JSONResponse(*clientError(e)), nil
	case not_found.Is(e):
		return server.DeleteMachineSnapshot404JSONResponse(*clientError(e)), nil
	case forbidden.Is(e):
		return server.DeleteMachineSnapshot403JSONResponse(*clientError(e)), nil
	case conflict.Is(e):
		return server.DeleteMachineSnapshot409JSONResponse(*clientError(e)), nil
	case unavailable.Is(e):
		return server.DeleteMachineSnapshot503JSONResponse(*clientError(e)), nil
	case e != nil:
		return server.DeleteMachineSnapshot500JSONResponse(*s.answer(e)), nil
	}

	return server.DeleteMachineSnapshot200JSONResponse{TaskId: taskIdentifier}, nil
}
