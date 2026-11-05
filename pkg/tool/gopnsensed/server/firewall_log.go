package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) FirewallLog(
	x context.Context,
	r server.FirewallLogRequestObject,
) (server.FirewallLogResponseObject, error) {
	var limit int

	if r.Params.Limit != nil {
		limit = *r.Params.Limit
	}

	result, e := s.service.Log(limit)

	if unavailable.Is(e) {
		return server.FirewallLog503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.FirewallLog500JSONResponse(*s.answer(x, e)), nil
	}

	return server.FirewallLog200JSONResponse(convert.LogEntries(result)), nil
}
