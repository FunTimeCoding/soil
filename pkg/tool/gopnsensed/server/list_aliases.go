package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ListAliases(
	_ context.Context,
	r server.ListAliasesRequestObject,
) (server.ListAliasesResponseObject, error) {
	var phrase string

	if r.Params.Query != nil {
		phrase = *r.Params.Query
	}

	result, e := s.opnsense.Aliases(phrase)

	if e != nil {
		return server.ListAliases500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListAliases200JSONResponse(convert.Aliases(result)), nil
}
