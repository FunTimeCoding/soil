package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ListRules(
	_ context.Context,
	r server.ListRulesRequestObject,
) (server.ListRulesResponseObject, error) {
	var phrase string

	if r.Params.Query != nil {
		phrase = *r.Params.Query
	}

	result, e := s.opnsense.Rules(phrase)

	if e != nil {
		return server.ListRules500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListRules200JSONResponse(convert.Rules(result)), nil
}
