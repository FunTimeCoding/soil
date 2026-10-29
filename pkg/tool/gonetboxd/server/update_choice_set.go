package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) UpdateChoiceSet(
	_ context.Context,
	r server.UpdateChoiceSetRequestObject,
) (server.UpdateChoiceSetResponseObject, error) {
	result, e := s.service.SetChoiceSetChoices(r.Name, r.Body.Choices)

	if not_found.Is(e) {
		return server.UpdateChoiceSet404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.UpdateChoiceSet400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.UpdateChoiceSet403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.UpdateChoiceSet500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.UpdateChoiceSet200JSONResponse(*convert.ChoiceSet(result)), nil
}
