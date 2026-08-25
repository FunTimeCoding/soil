package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ResetMachine(
	_ context.Context,
	r server.ResetMachineRequestObject,
) (server.ResetMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ResetMachine400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ResetMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	taskIdentifier, e := s.service.ResetMachine(c, int(r.Identifier), node)

	if e != nil {
		if not_found.Is(e) {
			return server.ResetMachine404JSONResponse{Error: e.Error()}, nil
		}

		return server.ResetMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ResetMachine200JSONResponse{TaskId: taskIdentifier}, nil
}
