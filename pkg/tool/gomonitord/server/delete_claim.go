package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	monitorConstant "github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/server"
)

func (s *Server) DeleteClaim(
	_ context.Context,
	r server.DeleteClaimRequestObject,
) (server.DeleteClaimResponseObject, error) {
	if r.Params.Item == "" || r.Params.Owner == "" {
		return server.DeleteClaim400JSONResponse{
			Error: monitorConstant.ClaimRequiredText,
		}, nil
	}

	if e := s.store.Release(r.Params.Item, r.Params.Owner); e != nil {
		return server.DeleteClaim500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.DeleteClaim204Response{}, nil
}
