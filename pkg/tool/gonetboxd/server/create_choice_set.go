package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateChoiceSet(
	x context.Context,
	r server.CreateChoiceSetRequestObject,
) (server.CreateChoiceSetResponseObject, error) {
	result, e := s.service.CreateChoiceSet(r.Body.Name, r.Body.Choices)

	if validation.Is(e) {
		return server.CreateChoiceSet400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreateChoiceSet503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateChoiceSet500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.CreateChoiceSet201JSONResponse(*convert.ChoiceSet(result)), nil
}
