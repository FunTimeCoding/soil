package server

import (
	"context"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetCheck(
	_ context.Context,
	r server.GetCheckRequestObject,
) (server.GetCheckResponseObject, error) {
	if r.Params.Preview != nil && *r.Params.Preview {
		return s.checkPreview(r.Params.Session)
	}

	result, checkError := s.service.Check(r.Params.Session)

	if checkError != nil {
		return server.GetCheck500JSONResponse(
			*s.captureFail(checkError, library.UnexpectedError),
		), nil
	}

	var entries []server.QueueEntry

	for _, e := range result.Entries {
		entries = append(
			entries,
			server.QueueEntry{
				Kind:      e.Kind,
				Body:      e.Body,
				Timestamp: e.CreatedAt.Format("2006-01-02T15:04:05Z"),
			},
		)
	}

	if entries == nil {
		entries = []server.QueueEntry{}
	}

	return server.GetCheck200JSONResponse(
		server.CheckResponse{
			Callsign: result.Callsign,
			Changed:  result.Changed,
			Entries:  entries,
		},
	), nil
}
