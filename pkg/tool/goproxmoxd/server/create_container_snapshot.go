package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) CreateContainerSnapshot(
	x context.Context,
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
		return server.CreateContainerSnapshot500JSONResponse(*s.answer(x, e)), nil
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

	switch {
	case validation.Is(e):
		return server.CreateContainerSnapshot400JSONResponse(*clientError(e)), nil
	case not_found.Is(e):
		return server.CreateContainerSnapshot404JSONResponse(*clientError(e)), nil
	case forbidden.Is(e):
		return server.CreateContainerSnapshot403JSONResponse(*clientError(e)), nil
	case conflict.Is(e):
		return server.CreateContainerSnapshot409JSONResponse(*clientError(e)), nil
	case unavailable.Is(e):
		return server.CreateContainerSnapshot503JSONResponse(*clientError(e)), nil
	case e != nil:
		return server.CreateContainerSnapshot500JSONResponse(*s.answer(x, e)), nil
	}

	return server.CreateContainerSnapshot200JSONResponse{TaskId: taskIdentifier}, nil
}
