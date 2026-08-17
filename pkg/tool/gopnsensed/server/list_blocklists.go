package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ListBlocklists(
	_ context.Context,
	r server.ListBlocklistsRequestObject,
) (server.ListBlocklistsResponseObject, error) {
	var phrase string

	if r.Params.Query != nil {
		phrase = *r.Params.Query
	}

	result, e := s.opnsense.Blocklists(phrase)

	if e != nil {
		return server.ListBlocklists500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListBlocklists200JSONResponse(convert.Blocklists(result)), nil
}
