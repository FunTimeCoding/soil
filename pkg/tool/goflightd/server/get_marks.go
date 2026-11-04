package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/server"
)

func (s *Server) GetMarks(
	x context.Context,
	r server.GetMarksRequestObject,
) (server.GetMarksResponseObject, error) {
	marks, e := s.service.Marks(limitOr(r.Params.Limit, 25))

	if e != nil {
		return server.GetMarks500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetMarks200JSONResponse(toMarks(marks)), nil
}
