package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/server"
)

func (s *Server) GetAlertsAfter(
	_ context.Context,
	r server.GetAlertsAfterRequestObject,
) (server.GetAlertsAfterResponseObject, error) {
	limit := constant.PageDefault

	if r.Params.Limit != nil {
		limit = min(max(*r.Params.Limit, 1), constant.PageMaximum)
	}

	records, e := s.service.AlertsAfter(
		uint(max(r.Params.Identifier, 0)),
		limit,
	)

	if e != nil {
		return server.GetAlertsAfter500JSONResponse(*s.answer(e)), nil
	}

	return server.GetAlertsAfter200JSONResponse(toResponse(records)), nil
}
