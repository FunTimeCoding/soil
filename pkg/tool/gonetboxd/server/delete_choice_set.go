package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) DeleteChoiceSet(
	_ context.Context,
	r server.DeleteChoiceSetRequestObject,
) (server.DeleteChoiceSetResponseObject, error) {
	e := s.service.DeleteChoiceSet(r.Name)

	if not_found.Is(e) {
		return server.DeleteChoiceSet404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.DeleteChoiceSet400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.DeleteChoiceSet403JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.DeleteChoiceSet409JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.DeleteChoiceSet500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteChoiceSet204Response{}, nil
}
