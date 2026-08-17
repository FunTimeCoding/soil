package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) FirewallStates(
	_ context.Context,
	r server.FirewallStatesRequestObject,
) (server.FirewallStatesResponseObject, error) {
	var phrase string

	if r.Params.Query != nil {
		phrase = *r.Params.Query
	}

	result, e := s.opnsense.States(phrase)

	if e != nil {
		return server.FirewallStates500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.FirewallStates200JSONResponse(convert.States(result)), nil
}
