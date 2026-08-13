package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) DeleteAddress(
	_ context.Context,
	r server.DeleteAddressRequestObject,
) (server.DeleteAddressResponseObject, error) {
	if e := s.client.DeleteInternet(r.Identifier); e != nil {
		return server.DeleteAddress500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteAddress204Response{}, nil
}
