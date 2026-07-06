package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) DeleteMachineSnapshot(
	_ context.Context,
	r server.DeleteMachineSnapshotRequestObject,
) (server.DeleteMachineSnapshotResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.DeleteMachineSnapshot400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.DeleteMachineSnapshot500JSONResponse(*s.captureDetail(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	taskID, e := s.service.DeleteMachineSnapshot(
		c,
		int(r.Identifier),
		node,
		r.Name,
	)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return server.DeleteMachineSnapshot404JSONResponse{Error: e.Error()}, nil
		}

		return server.DeleteMachineSnapshot500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteMachineSnapshot200JSONResponse{TaskId: taskID}, nil
}
