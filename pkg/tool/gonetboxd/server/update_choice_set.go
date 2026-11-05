package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) UpdateChoiceSet(
	x context.Context,
	r server.UpdateChoiceSetRequestObject,
) (server.UpdateChoiceSetResponseObject, error) {
	result, e := s.service.SetChoiceSetChoices(r.Name, r.Body.Choices)

	if not_found.Is(e) {
		return server.UpdateChoiceSet404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.UpdateChoiceSet400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.UpdateChoiceSet503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.UpdateChoiceSet500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.UpdateChoiceSet200JSONResponse(*convert.ChoiceSet(result)), nil
}
