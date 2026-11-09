package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/server"
)

func (s *Server) GetTopAlerts(
	x context.Context,
	r server.GetTopAlertsRequestObject,
) (server.GetTopAlertsResponseObject, error) {
	n := 25
	w := window.Whole()

	if r.Params.N != nil {
		n = *r.Params.N
	}

	if r.Params.Start != nil {
		w.WithStart(*r.Params.Start)
	}

	if r.Params.End != nil {
		w.WithEnd(*r.Params.End)
	}

	records, e := s.service.TopAlerts(n, w)

	if e != nil {
		return server.GetTopAlerts500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetTopAlerts200JSONResponse(convert.TopAlerts(records)), nil
}
