package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) StopMachine(
	_ context.Context,
	r server.StopMachineRequestObject,
) (server.StopMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.StopMachine400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.StopMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	taskID, e := s.service.StopMachine(c, int(r.Identifier), node)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return server.StopMachine404JSONResponse{Error: e.Error()}, nil
		}

		return server.StopMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.StopMachine200JSONResponse{TaskId: taskID}, nil
}
