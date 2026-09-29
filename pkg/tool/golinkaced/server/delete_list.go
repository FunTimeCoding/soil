package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) DeleteList(
	_ context.Context,
	r server.DeleteListRequestObject,
) (server.DeleteListResponseObject, error) {
	if e := s.client.DeleteList(int(r.Identifier)); e != nil {
		return server.DeleteList500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteList200Response{}, nil
}
