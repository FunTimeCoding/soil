package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetChannel(
	x context.Context,
	r server.GetChannelRequestObject,
) (server.GetChannelResponseObject, error) {
	holder, e := s.service.SessionByCallsign(r.Params.Callsign)

	if e != nil {
		return server.GetChannel500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	if holder == nil {
		return server.GetChannel200JSONResponse(
			server.CheckResponse{Entries: []server.QueueEntry{}},
		), nil
	}

	drained, f := s.service.AwaitImmediateQueue(
		x,
		holder.Identifier,
		r.Params.Callsign,
	)

	if f != nil {
		return server.GetChannel500JSONResponse(
			*s.captureFail(f, constant.UnexpectedError),
		), nil
	}

	entries := []server.QueueEntry{}

	for _, entry := range drained {
		entries = append(
			entries,
			server.QueueEntry{
				Kind:      entry.Kind,
				Body:      entry.Body,
				Timestamp: entry.CreatedAt.Format("2006-01-02T15:04:05Z"),
			},
		)
	}

	return server.GetChannel200JSONResponse(
		server.CheckResponse{
			Callsign: r.Params.Callsign,
			Changed:  len(entries) > 0,
			Entries:  entries,
		},
	), nil
}
