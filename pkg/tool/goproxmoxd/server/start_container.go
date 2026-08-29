package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) StartContainer(
	_ context.Context,
	r server.StartContainerRequestObject,
) (server.StartContainerResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.StartContainer400JSONResponse(*clientError(e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	taskIdentifier, e := s.service.StartContainer(
		instance,
		int(r.Identifier),
		node,
	)

	if e != nil {
		if not_found.Is(e) {
			return server.StartContainer404JSONResponse{Error: e.Error()}, nil
		}

		return server.StartContainer500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.StartContainer200JSONResponse{TaskId: taskIdentifier}, nil
}
