package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListDeviceJournalEntries(
	_ context.Context,
	r server.ListDeviceJournalEntriesRequestObject,
) (server.ListDeviceJournalEntriesResponseObject, error) {
	var limit int32
	var offset int32

	if r.Params.Limit != nil {
		limit = *r.Params.Limit
	}

	if r.Params.Offset != nil {
		offset = *r.Params.Offset
	}

	result, e := s.client.DeviceJournalEntries(r.Name, limit, offset)

	if not_found.Is(e) {
		return server.ListDeviceJournalEntries404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListDeviceJournalEntries500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.ListDeviceJournalEntries200JSONResponse(
		convert.JournalEntries(result),
	), nil
}
