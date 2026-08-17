package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ListHosts(
	_ context.Context,
	r server.ListHostsRequestObject,
) (server.ListHostsResponseObject, error) {
	var phrase string

	if r.Params.Query != nil {
		phrase = *r.Params.Query
	}

	result, e := s.opnsense.Hosts(phrase)

	if e != nil {
		return server.ListHosts500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListHosts200JSONResponse(convert.Hosts(result)), nil
}
