package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ListHosts(
	x context.Context,
	r server.ListHostsRequestObject,
) (server.ListHostsResponseObject, error) {
	var phrase string

	if r.Params.Query != nil {
		phrase = *r.Params.Query
	}

	result, e := s.service.Hosts(phrase)

	if unavailable.Is(e) {
		return server.ListHosts503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListHosts500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ListHosts200JSONResponse(convert.Hosts(result)), nil
}
