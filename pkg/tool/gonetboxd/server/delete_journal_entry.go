package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) DeleteJournalEntry(
	x context.Context,
	r server.DeleteJournalEntryRequestObject,
) (server.DeleteJournalEntryResponseObject, error) {
	e := s.service.DeleteJournalEntry(r.Identifier)

	if validation.Is(e) {
		return server.DeleteJournalEntry400JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.DeleteJournalEntry409JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.DeleteJournalEntry503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteJournalEntry500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.DeleteJournalEntry204Response{}, nil
}
