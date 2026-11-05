package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateCustomField(
	x context.Context,
	r server.CreateCustomFieldRequestObject,
) (server.CreateCustomFieldResponseObject, error) {
	var label string

	if r.Body.Label != nil {
		label = *r.Body.Label
	}

	result, e := s.service.CreateSelectionField(
		r.Body.Name,
		label,
		r.Body.ChoiceSet,
	)

	if not_found.Is(e) {
		return server.CreateCustomField404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.CreateCustomField400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreateCustomField503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateCustomField500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.CreateCustomField201JSONResponse(*convert.CustomField(result)),
		nil
}
