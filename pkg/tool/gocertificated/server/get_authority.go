package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) GetAuthority(
	_ context.Context,
	r server.GetAuthorityRequestObject,
) (server.GetAuthorityResponseObject, error) {
	result, e := s.store.Authority(r.Name)

	if e != nil {
		return server.GetAuthority500JSONResponse(
			*s.captureFail(e, constant.QueryFail),
		), nil
	}

	if result == nil {
		return server.GetAuthority404JSONResponse(
			*clientError(constant.AuthorityMissing),
		), nil
	}

	return server.GetAuthority200JSONResponse(*convert.Authority(result)), nil
}
