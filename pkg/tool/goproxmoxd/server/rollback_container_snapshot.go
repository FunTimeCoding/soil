package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) RollbackContainerSnapshot(
	_ context.Context,
	r server.RollbackContainerSnapshotRequestObject,
) (server.RollbackContainerSnapshotResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.RollbackContainerSnapshot400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.RollbackContainerSnapshot500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	taskID, e := s.service.RollbackContainerSnapshot(
		c,
		int(r.Identifier),
		node,
		r.Name,
	)

	if e != nil {
		if not_found.Is(e) {
			return server.RollbackContainerSnapshot404JSONResponse{
				Error: e.Error(),
			}, nil
		}

		return server.RollbackContainerSnapshot500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.RollbackContainerSnapshot200JSONResponse{TaskId: taskID}, nil
}
