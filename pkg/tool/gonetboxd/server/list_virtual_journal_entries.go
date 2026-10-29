package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListVirtualJournalEntries(
	_ context.Context,
	r server.ListVirtualJournalEntriesRequestObject,
) (server.ListVirtualJournalEntriesResponseObject, error) {
	var limit int32
	var offset int32

	if r.Params.Limit != nil {
		limit = *r.Params.Limit
	}

	if r.Params.Offset != nil {
		offset = *r.Params.Offset
	}

	result, e := s.service.VirtualJournalEntries(r.Name, limit, offset)

	if not_found.Is(e) {
		return server.ListVirtualJournalEntries404JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.ListVirtualJournalEntries403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListVirtualJournalEntries500JSONResponse(
				*s.captureDetail(e),
			),
			nil
	}

	return server.ListVirtualJournalEntries200JSONResponse(
			convert.JournalEntries(result),
		),
		nil
}
