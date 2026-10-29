package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
)

func (s *Server) PostEntry(
	_ context.Context,
	r server.PostEntryRequestObject,
) (server.PostEntryResponseObject, error) {
	e := entry.New()
	e.Action = r.Body.Action
	e.User = r.Body.User

	if r.Body.System != nil {
		e.System = *r.Body.System
	}

	if r.Body.Service != nil {
		e.Service = *r.Body.Service
	}

	if r.Body.Description != nil {
		e.Description = *r.Body.Description
	}

	if r.Body.Timestamp != nil {
		e.Timestamp = *r.Body.Timestamp
	}

	f := s.service.Add(e)

	if validation.Is(f) {
		return server.PostEntry400JSONResponse{Error: f.Error()}, nil
	}

	if f != nil {
		return server.PostEntry500JSONResponse(*s.answer(f)), nil
	}

	return server.PostEntry200JSONResponse(toResponse([]entry.Entry{*e})[0]), nil
}
