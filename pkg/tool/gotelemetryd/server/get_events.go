package server

import (
	"context"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
)

func (s *Server) GetEvents(
	_ context.Context,
	r server.GetEventsRequestObject,
) (server.GetEventsResponseObject, error) {
	o := store.NewQueryOption()

	if r.Params.Tool != nil {
		o.Tool = *r.Params.Tool
	}

	if r.Params.Surface != nil {
		o.Surface = *r.Params.Surface
	}

	if r.Params.Actor != nil {
		o.Actor = *r.Params.Actor
	}

	if r.Params.Kind != nil {
		o.Kind = *r.Params.Kind
	}

	if r.Params.Since != nil {
		o.Since = *r.Params.Since
	}

	if r.Params.Until != nil {
		o.Until = *r.Params.Until
	}

	if r.Params.Limit != nil {
		o.Limit = *r.Params.Limit
	}

	if r.Params.Offset != nil {
		o.Offset = *r.Params.Offset
	}

	events, e := s.store.Recent(o)

	if e != nil {
		return server.GetEvents500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	entries := make([]server.EventEntry, 0, len(events))

	for _, e := range events {
		entry := server.EventEntry{
			Id:        int(e.ID),
			Tool:      e.Tool,
			Surface:   e.Surface,
			Actor:     e.Actor,
			Outcome:   e.Outcome,
			Kind:      e.Kind,
			CreatedAt: e.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}

		if e.DurationMillisecond > 0 {
			entry.DurationMs = new(int(e.DurationMillisecond))
		}

		if e.Detail != nil {
			var detail map[string]string

			if json.Unmarshal([]byte(*e.Detail), &detail) == nil {
				entry.Detail = &detail
			}
		}

		entries = append(entries, entry)
	}

	return server.GetEvents200JSONResponse(entries), nil
}
