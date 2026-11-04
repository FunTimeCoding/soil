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

func (s *Server) StopContainer(
	x context.Context,
	r server.StopContainerRequestObject,
) (server.StopContainerResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.StopContainer404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.StopContainer400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.StopContainer500JSONResponse(*s.answer(x, e)), nil
	}

	node := ""

	if r.Params.Node != nil {
		node = *r.Params.Node
	}

	taskIdentifier, e := s.service.StopContainer(
		instance,
		int(r.Identifier),
		node,
	)

	switch {
	case validation.Is(e):
		return server.StopContainer400JSONResponse(*clientError(e)), nil
	case not_found.Is(e):
		return server.StopContainer404JSONResponse(*clientError(e)), nil
	case forbidden.Is(e):
		return server.StopContainer403JSONResponse(*clientError(e)), nil
	case conflict.Is(e):
		return server.StopContainer409JSONResponse(*clientError(e)), nil
	case unavailable.Is(e):
		return server.StopContainer503JSONResponse(*clientError(e)), nil
	case e != nil:
		return server.StopContainer500JSONResponse(*s.answer(x, e)), nil
	}

	return server.StopContainer200JSONResponse{TaskId: taskIdentifier}, nil
}
