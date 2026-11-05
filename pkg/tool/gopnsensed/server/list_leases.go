package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ListLeases(
	x context.Context,
	r server.ListLeasesRequestObject,
) (server.ListLeasesResponseObject, error) {
	var phrase string

	if r.Params.Query != nil {
		phrase = *r.Params.Query
	}

	result, e := s.service.Leases(phrase)

	if unavailable.Is(e) {
		return server.ListLeases503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListLeases500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ListLeases200JSONResponse(convert.Leases(result)), nil
}
