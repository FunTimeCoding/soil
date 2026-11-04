package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/server"
)

func (s *Server) GetEvents(
	x context.Context,
	r server.GetEventsRequestObject,
) (server.GetEventsResponseObject, error) {
	start, end := timeRange(r.Params.Start, r.Params.End)
	events, e := s.service.Events(start, end, limitOr(r.Params.Limit, 1000))

	if e != nil {
		return server.GetEvents500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetEvents200JSONResponse(toEvents(events)), nil
}
