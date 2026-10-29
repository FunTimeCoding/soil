package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) UpdateJournalEntry(
	_ context.Context,
	r server.UpdateJournalEntryRequestObject,
) (server.UpdateJournalEntryResponseObject, error) {
	var kind string
	var comments string

	if r.Body.Kind != nil {
		kind = *r.Body.Kind
	}

	if r.Body.Comments != nil {
		comments = *r.Body.Comments
	}

	result, e := s.service.UpdateJournalEntry(r.Identifier, kind, comments)

	if validation.Is(e) {
		return server.UpdateJournalEntry400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.UpdateJournalEntry403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.UpdateJournalEntry500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.UpdateJournalEntry200JSONResponse(
			*convert.JournalEntry(result),
		),
		nil
}
