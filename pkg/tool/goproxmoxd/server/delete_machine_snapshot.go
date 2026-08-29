package server

import (
	"context"
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

	if e != nil {
		if not_found.Is(e) {
			return server.DeleteMachineSnapshot404JSONResponse{Error: e.Error()}, nil
		}

		return server.DeleteMachineSnapshot500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteMachineSnapshot200JSONResponse{TaskId: taskIdentifier}, nil
}
