package server

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) checkPreview(
	sessionIdentifier string,
) (server.GetCheckResponseObject, error) {
	callsign, e := s.roster.CallsignBySessionIdentifier(sessionIdentifier)

	if e != nil {
		return server.GetCheck500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	if callsign == "" {
		return server.GetCheck200JSONResponse{
			Changed: false,
			Entries: []server.QueueEntry{},
		}, nil
	}

	entries, e := s.message.PeekQueue(sessionIdentifier, callsign)

	if e != nil {
		return server.GetCheck500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	text, e := s.message.Render(entries)

	if e != nil {
		return server.GetCheck500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	var result []server.QueueEntry

	for _, entry := range entries {
		result = append(
			result,
			server.QueueEntry{
				Kind:      entry.Kind,
				Body:      entry.Body,
				Timestamp: entry.CreatedAt.Format("2006-01-02T15:04:05Z"),
			},
		)
	}

	if result == nil {
		result = []server.QueueEntry{}
	}

	return server.GetCheck200JSONResponse{
		Callsign: callsign,
		Changed:  len(result) > 0,
		Context:  text,
		Entries:  result,
	}, nil
}
