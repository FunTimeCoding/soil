package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) GetRevocationList(
	_ context.Context,
	r server.GetRevocationListRequestObject,
) (server.GetRevocationListResponseObject, error) {
	result, e := s.service.RevocationList(r.Params.Authority)

	if not_found.Is(e) {
		return server.GetRevocationList404JSONResponse(
			*clientError(constant.AuthorityMissing),
		), nil
	}

	if e != nil {
		return server.GetRevocationList500JSONResponse(
			*s.captureFail(e, constant.RevocationListFail),
		), nil
	}

	return server.GetRevocationList200TextResponse(result), nil
}
