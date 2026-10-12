package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/server"
)

func (s *Server) DeleteClaim(
	_ context.Context,
	r server.DeleteClaimRequestObject,
) (server.DeleteClaimResponseObject, error) {
	if r.Params.Item == "" || r.Params.Owner == "" {
		return server.DeleteClaim400JSONResponse{
			Error: constant.ClaimRequiredText,
		}, nil
	}

	if e := s.store.Release(r.Params.Item, r.Params.Owner); e != nil {
		return server.DeleteClaim500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	return server.DeleteClaim204Response{}, nil
}
