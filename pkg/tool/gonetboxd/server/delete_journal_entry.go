package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) DeleteJournalEntry(
	_ context.Context,
	r server.DeleteJournalEntryRequestObject,
) (server.DeleteJournalEntryResponseObject, error) {
	if e := s.client.DeleteJournalEntry(r.Identifier); e != nil {
		return server.DeleteJournalEntry500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.DeleteJournalEntry204Response{}, nil
}
