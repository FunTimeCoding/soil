package server

import (
	"context"
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

	result, e := s.client.AddVirtualJournalEntry(
		r.Name,
		kind,
		r.Body.Comments,
	)

	if e != nil {
		return server.AddVirtualJournalEntry500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.AddVirtualJournalEntry201JSONResponse(
		*convert.JournalEntry(result),
	), nil
}
