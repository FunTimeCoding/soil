package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/server"
)

func (s *Server) GetAlerts(
	_ context.Context,
	r server.GetAlertsRequestObject,
) (server.GetAlertsResponseObject, error) {
	records, e := s.store.ByName(r.Params.Name)

	if e != nil {
		return server.GetAlerts500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.GetAlerts200JSONResponse(toResponse(records)), nil
}
