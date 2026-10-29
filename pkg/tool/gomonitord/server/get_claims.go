package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/convert"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/server"
)

func (s *Server) GetClaims(
	_ context.Context,
	_ server.GetClaimsRequestObject,
) (server.GetClaimsResponseObject, error) {
	v, e := s.service.Claims()

	if e != nil {
		return server.GetClaims500JSONResponse(*s.answer(e)), nil
	}

	return server.GetClaims200JSONResponse(convert.Claims(v)), nil
}
