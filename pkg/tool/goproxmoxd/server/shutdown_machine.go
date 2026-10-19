package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ShutdownMachine(
	_ context.Context,
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
		return server.ShutdownMachine500JSONResponse(
			*s.captureFail(e, e.Error()),
		), nil
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

	if e != nil {
		if not_found.Is(e) {
			return server.ShutdownMachine404JSONResponse{Error: e.Error()}, nil
		}

		return server.ShutdownMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ShutdownMachine200JSONResponse{TaskId: taskIdentifier}, nil
}
