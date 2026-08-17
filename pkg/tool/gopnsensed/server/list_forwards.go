package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ListForwards(
	_ context.Context,
	r server.ListForwardsRequestObject,
) (server.ListForwardsResponseObject, error) {
	var phrase string

	if r.Params.Query != nil {
		phrase = *r.Params.Query
	}

	result, e := s.opnsense.Forwards(phrase)

	if e != nil {
		return server.ListForwards500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListForwards200JSONResponse(convert.Forwards(result)), nil
}
