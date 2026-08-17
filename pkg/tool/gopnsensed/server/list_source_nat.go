package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ListSourceNat(
	_ context.Context,
	r server.ListSourceNatRequestObject,
) (server.ListSourceNatResponseObject, error) {
	var phrase string

	if r.Params.Query != nil {
		phrase = *r.Params.Query
	}

	result, e := s.opnsense.SourceNatRules(phrase)

	if e != nil {
		return server.ListSourceNat500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListSourceNat200JSONResponse(
		convert.SourceNatRules(result),
	), nil
}
