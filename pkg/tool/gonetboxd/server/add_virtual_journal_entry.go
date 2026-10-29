package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) AddVirtualJournalEntry(
	_ context.Context,
	r server.AddVirtualJournalEntryRequestObject,
) (server.AddVirtualJournalEntryResponseObject, error) {
	var kind string

	if r.Body.Kind != nil {
		kind = *r.Body.Kind
	}

	result, e := s.service.AddVirtualJournalEntry(r.Name, kind, r.Body.Comments)

	if not_found.Is(e) {
		return server.AddVirtualJournalEntry404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.AddVirtualJournalEntry400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.AddVirtualJournalEntry403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.AddVirtualJournalEntry500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.AddVirtualJournalEntry201JSONResponse(
			*convert.JournalEntry(result),
		),
		nil
}
