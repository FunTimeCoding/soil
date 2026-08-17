package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ListPools(
	_ context.Context,
	r server.ListPoolsRequestObject,
) (server.ListPoolsResponseObject, error) {
	var phrase string

	if r.Params.Query != nil {
		phrase = *r.Params.Query
	}

	result, e := s.opnsense.Pools(phrase)

	if e != nil {
		return server.ListPools500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListPools200JSONResponse(convert.Pools(result)), nil
}
