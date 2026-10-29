package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/server"
)

func (s *Server) PostMark(
	_ context.Context,
	r server.PostMarkRequestObject,
) (server.PostMarkResponseObject, error) {
	var label string
	var note string

	if r.Body != nil {
		label = r.Body.Label

		if r.Body.Note != nil {
			note = *r.Body.Note
		}
	}

	v, e := s.service.CreateMark(label, note)

	if validation.Is(e) {
		return server.PostMark400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PostMark500JSONResponse(*s.answer(e)), nil
	}

	return server.PostMark201JSONResponse(*toMark(*v)), nil
}
