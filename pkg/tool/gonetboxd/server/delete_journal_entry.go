package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) DeleteJournalEntry(
	_ context.Context,
	r server.DeleteJournalEntryRequestObject,
) (server.DeleteJournalEntryResponseObject, error) {
	e := s.service.DeleteJournalEntry(r.Identifier)

	if validation.Is(e) {
		return server.DeleteJournalEntry400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.DeleteJournalEntry403JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.DeleteJournalEntry409JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.DeleteJournalEntry500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteJournalEntry204Response{}, nil
}
