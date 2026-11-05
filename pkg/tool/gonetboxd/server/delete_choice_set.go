package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) DeleteChoiceSet(
	x context.Context,
	r server.DeleteChoiceSetRequestObject,
) (server.DeleteChoiceSetResponseObject, error) {
	e := s.service.DeleteChoiceSet(r.Name)

	if not_found.Is(e) {
		return server.DeleteChoiceSet404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.DeleteChoiceSet400JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.DeleteChoiceSet409JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.DeleteChoiceSet503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteChoiceSet500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.DeleteChoiceSet204Response{}, nil
}
