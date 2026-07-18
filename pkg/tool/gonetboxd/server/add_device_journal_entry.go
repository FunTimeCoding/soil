package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) AddDeviceJournalEntry(
	_ context.Context,
	r server.AddDeviceJournalEntryRequestObject,
) (server.AddDeviceJournalEntryResponseObject, error) {
	var kind string

	if r.Body.Kind != nil {
		kind = *r.Body.Kind
	}

	result, e := s.client.AddDeviceJournalEntry(
		r.Name,
		kind,
		r.Body.Comments,
	)

	if e != nil {
		return server.AddDeviceJournalEntry500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.AddDeviceJournalEntry201JSONResponse(
		*convert.JournalEntry(result),
	), nil
}
