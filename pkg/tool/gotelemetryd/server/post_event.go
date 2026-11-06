package server

import (
	"context"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model/usage_event"
)

func (s *Server) PostEvent(
	x context.Context,
	r server.PostEventRequestObject,
) (server.PostEventResponseObject, error) {
	e := usage_event.Stub()
	e.Tool = r.Body.Tool
	e.Surface = r.Body.Surface
	e.Actor = r.Body.Actor
	e.Outcome = r.Body.Outcome
	e.Kind = r.Body.Kind

	if r.Body.DurationMs != nil {
		e.DurationMillisecond = int64(*r.Body.DurationMs)
	}

	if r.Body.Detail != nil {
		encoded, marshalError := json.Marshal(*r.Body.Detail)

		if marshalError == nil {
			e.Detail = new(string(encoded))
		}
	}

	if f := s.service.Create(e); f != nil {
		return server.PostEvent500JSONResponse(*s.answer(x, f)), nil
	}

	return server.PostEvent200JSONResponse{Id: int(e.Identifier)}, nil
}
