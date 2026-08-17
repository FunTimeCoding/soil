package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ListLeases(
	_ context.Context,
	r server.ListLeasesRequestObject,
) (server.ListLeasesResponseObject, error) {
	var phrase string

	if r.Params.Query != nil {
		phrase = *r.Params.Query
	}

	result, e := s.opnsense.Leases(phrase)

	if e != nil {
		return server.ListLeases500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListLeases200JSONResponse(convert.Leases(result)), nil
}
