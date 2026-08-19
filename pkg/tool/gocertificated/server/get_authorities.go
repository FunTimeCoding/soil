package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) GetAuthorities(
	_ context.Context,
	_ server.GetAuthoritiesRequestObject,
) (server.GetAuthoritiesResponseObject, error) {
	result, e := s.store.Authorities()

	if e != nil {
		return server.GetAuthorities500JSONResponse(
			*s.captureFail(e, constant.QueryFail),
		), nil
	}

	return server.GetAuthorities200JSONResponse(
		convert.Authorities(result),
	), nil
}
