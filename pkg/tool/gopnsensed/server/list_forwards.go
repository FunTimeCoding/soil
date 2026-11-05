package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ListForwards(
	x context.Context,
	r server.ListForwardsRequestObject,
) (server.ListForwardsResponseObject, error) {
	var phrase string

	if r.Params.Query != nil {
		phrase = *r.Params.Query
	}

	result, e := s.service.Forwards(phrase)

	if unavailable.Is(e) {
		return server.ListForwards503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListForwards500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ListForwards200JSONResponse(convert.Forwards(result)), nil
}
