package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/server"
)

func (s *Server) GetAlerts(
	x context.Context,
	r server.GetAlertsRequestObject,
) (server.GetAlertsResponseObject, error) {
	records, e := s.service.Alerts(r.Params.Name)

	if e != nil {
		return server.GetAlerts500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetAlerts200JSONResponse(toResponse(records)), nil
}
