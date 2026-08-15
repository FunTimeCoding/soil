package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/server"
)

func (s *Server) GetMarks(
	_ context.Context,
	r server.GetMarksRequestObject,
) (server.GetMarksResponseObject, error) {
	marks, e := s.store.RecentMarks(limitOr(r.Params.Limit, 25))

	if e != nil {
		return server.GetMarks500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.GetMarks200JSONResponse(toMarks(marks)), nil
}
