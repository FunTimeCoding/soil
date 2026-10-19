package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) CreateContainerSnapshot(
	_ context.Context,
	r server.CreateContainerSnapshotRequestObject,
) (server.CreateContainerSnapshotResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.CreateContainerSnapshot404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.CreateContainerSnapshot400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreateContainerSnapshot500JSONResponse(
			*s.captureFail(e, e.Error()),
		), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	taskIdentifier, e := s.service.CreateContainerSnapshot(
		instance,
		int(r.Identifier),
		node,
		r.Body.Name,
	)

	if e != nil {
		if not_found.Is(e) {
			return server.CreateContainerSnapshot404JSONResponse{
				Error: e.Error(),
			}, nil
		}

		return server.CreateContainerSnapshot500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.CreateContainerSnapshot200JSONResponse{TaskId: taskIdentifier}, nil
}
