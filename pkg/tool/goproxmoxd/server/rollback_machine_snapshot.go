package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) RollbackMachineSnapshot(
	_ context.Context,
	r server.RollbackMachineSnapshotRequestObject,
) (server.RollbackMachineSnapshotResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.RollbackMachineSnapshot400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.RollbackMachineSnapshot500JSONResponse(*s.captureDetail(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	taskID, e := s.service.RollbackMachineSnapshot(
		c,
		int(r.Identifier),
		node,
		r.Name,
	)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return server.RollbackMachineSnapshot404JSONResponse{Error: e.Error()}, nil
		}

		return server.RollbackMachineSnapshot500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.RollbackMachineSnapshot200JSONResponse{TaskId: taskID}, nil
}
