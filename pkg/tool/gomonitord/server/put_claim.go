package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/convert"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/server"
)

func (s *Server) PutClaim(
	_ context.Context,
	r server.PutClaimRequestObject,
) (server.PutClaimResponseObject, error) {
	c, e := s.service.ClaimItem(r.Body.Item, r.Body.Owner)

	if validation.Is(e) {
		return server.PutClaim400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PutClaim500JSONResponse(*s.answer(e)), nil
	}

	return server.PutClaim200JSONResponse(convert.Claim(c)), nil
}
