package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ListSourceNat(
	x context.Context,
	r server.ListSourceNatRequestObject,
) (server.ListSourceNatResponseObject, error) {
	var phrase string

	if r.Params.Query != nil {
		phrase = *r.Params.Query
	}

	result, e := s.service.SourceNatRules(phrase)

	if unavailable.Is(e) {
		return server.ListSourceNat503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListSourceNat500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ListSourceNat200JSONResponse(
		convert.SourceNatRules(result),
	), nil
}
