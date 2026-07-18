package server

import (
	"context"
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

	result, e := s.client.UpdateJournalEntry(
		r.Identifier,
		kind,
		comments,
	)

	if e != nil {
		return server.UpdateJournalEntry500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.UpdateJournalEntry200JSONResponse(
		*convert.JournalEntry(result),
	), nil
}
