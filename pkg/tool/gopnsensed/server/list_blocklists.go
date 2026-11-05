package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ListBlocklists(
	x context.Context,
	r server.ListBlocklistsRequestObject,
) (server.ListBlocklistsResponseObject, error) {
	var phrase string

	if r.Params.Query != nil {
		phrase = *r.Params.Query
	}

	result, e := s.service.Blocklists(phrase)

	if unavailable.Is(e) {
		return server.ListBlocklists503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListBlocklists500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ListBlocklists200JSONResponse(convert.Blocklists(result)), nil
}
