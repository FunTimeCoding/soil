package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) PostAuthority(
	_ context.Context,
	r server.PostAuthorityRequestObject,
) (server.PostAuthorityResponseObject, error) {
	result, e := s.service.CreateAuthority(r.Body)

	if conflict.Is(e) {
		return server.PostAuthority409JSONResponse(
			*clientError(constant.AuthorityLive),
		), nil
	}

	if e != nil {
		return server.PostAuthority500JSONResponse(
			*s.captureFail(e, constant.CreateAuthorityFail),
		), nil
	}

	return server.PostAuthority200JSONResponse(*convert.Authority(result)), nil
}
