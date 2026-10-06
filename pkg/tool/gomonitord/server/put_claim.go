package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	monitorConstant "github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/convert"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/server"
)

func (s *Server) PutClaim(
	_ context.Context,
	r server.PutClaimRequestObject,
) (server.PutClaimResponseObject, error) {
	if r.Body.Item == "" || r.Body.Owner == "" {
		return server.PutClaim400JSONResponse{
			Error: monitorConstant.ClaimRequiredText,
		}, nil
	}

	c, e := s.store.ClaimItem(r.Body.Item, r.Body.Owner)

	if e != nil {
		return server.PutClaim500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PutClaim200JSONResponse(convert.Claim(c)), nil
}
