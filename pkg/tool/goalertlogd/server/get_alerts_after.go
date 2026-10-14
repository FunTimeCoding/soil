package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
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

	records, e := s.store.After(uint(max(r.Params.Identifier, 0)), limit)

	if e != nil {
		return server.GetAlertsAfter500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	return server.GetAlertsAfter200JSONResponse(toResponse(records)), nil
}
