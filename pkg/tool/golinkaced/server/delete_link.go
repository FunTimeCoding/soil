package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) DeleteLink(
	_ context.Context,
	r server.DeleteLinkRequestObject,
) (server.DeleteLinkResponseObject, error) {
	e := s.client.DeleteLink(int(r.Identifier))

	if not_found.Is(e) {
		return server.DeleteLink404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteLink500JSONResponse(*s.answer(e)), nil
	}

	return server.DeleteLink200Response{}, nil
}
