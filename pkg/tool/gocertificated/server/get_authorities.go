package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) GetAuthorities(
	x context.Context,
	_ server.GetAuthoritiesRequestObject,
) (server.GetAuthoritiesResponseObject, error) {
	result, e := s.service.Authorities()

	if e != nil {
		return server.GetAuthorities500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetAuthorities200JSONResponse(
		convert.Authorities(result),
	), nil
}
