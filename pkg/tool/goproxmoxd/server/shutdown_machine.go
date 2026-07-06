package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ShutdownMachine(
	_ context.Context,
	r server.ShutdownMachineRequestObject,
) (server.ShutdownMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ShutdownMachine400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ShutdownMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	taskID, e := s.service.ShutdownMachine(c, int(r.Identifier), node)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return server.ShutdownMachine404JSONResponse{Error: e.Error()}, nil
		}

		return server.ShutdownMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ShutdownMachine200JSONResponse{TaskId: taskID}, nil
}
