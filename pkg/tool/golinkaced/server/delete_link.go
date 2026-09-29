package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) DeleteLink(
	_ context.Context,
	r server.DeleteLinkRequestObject,
) (server.DeleteLinkResponseObject, error) {
	if e := s.client.DeleteLink(int(r.Identifier)); e != nil {
		return server.DeleteLink500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteLink200Response{}, nil
}
