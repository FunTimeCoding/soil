package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) DeleteContainerSnapshot(
	_ context.Context,
	r server.DeleteContainerSnapshotRequestObject,
) (server.DeleteContainerSnapshotResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.DeleteContainerSnapshot400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.DeleteContainerSnapshot500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	taskID, e := s.service.DeleteContainerSnapshot(
		c,
		int(r.Identifier),
		node,
		r.Name,
	)

	if e != nil {
		if not_found.Is(e) {
			return server.DeleteContainerSnapshot404JSONResponse{
				Error: e.Error(),
			}, nil
		}

		return server.DeleteContainerSnapshot500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.DeleteContainerSnapshot200JSONResponse{TaskId: taskID}, nil
}
