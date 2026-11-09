package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/server"
)

func (s *Server) GetRecentAlerts(
	x context.Context,
	r server.GetRecentAlertsRequestObject,
) (server.GetRecentAlertsResponseObject, error) {
	w := window.Whole()

	if r.Params.Start != nil {
		w.WithStart(*r.Params.Start)
	}

	if r.Params.End != nil {
		w.WithEnd(*r.Params.End)
	}

	records, e := s.service.RecentAlerts(w)

	if e != nil {
		return server.GetRecentAlerts500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetRecentAlerts200JSONResponse(toResponse(records)), nil
}
