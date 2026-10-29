package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateCustomField(
	_ context.Context,
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

	if forbidden.Is(e) {
		return server.CreateCustomField403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreateCustomField500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateCustomField201JSONResponse(*convert.CustomField(result)),
		nil
}
