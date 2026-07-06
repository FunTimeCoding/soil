package server

import (
	"context"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
)

func (s *Server) PostEvent(
	_ context.Context,
	r server.PostEventRequestObject,
) (server.PostEventResponseObject, error) {
	e := store.NewUsageEvent()
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

	if f := s.store.Create(e); f != nil {
		return server.PostEvent500JSONResponse(
			*s.captureFail(f, constant.UnexpectedError),
		), nil
	}

	return server.PostEvent200JSONResponse{Id: int(e.ID)}, nil
}
