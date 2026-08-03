package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) StartMachine(
	_ context.Context,
	r server.StartMachineRequestObject,
) (server.StartMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.StartMachine400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.StartMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	taskID, e := s.service.StartMachine(c, int(r.Identifier), node)

	if e != nil {
		if not_found.Is(e) {
			return server.StartMachine404JSONResponse{Error: e.Error()}, nil
		}

		return server.StartMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.StartMachine200JSONResponse{TaskId: taskID}, nil
}
