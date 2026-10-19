package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) StartMachine(
	_ context.Context,
	r server.StartMachineRequestObject,
) (server.StartMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.StartMachine404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.StartMachine400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.StartMachine500JSONResponse(*s.captureFail(e, e.Error())), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	taskIdentifier, e := s.service.StartMachine(
		instance,
		int(r.Identifier),
		node,
	)

	if e != nil {
		if not_found.Is(e) {
			return server.StartMachine404JSONResponse{Error: e.Error()}, nil
		}

		return server.StartMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.StartMachine200JSONResponse{TaskId: taskIdentifier}, nil
}
