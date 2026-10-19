package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) DeleteTag(
	_ context.Context,
	r server.DeleteTagRequestObject,
) (server.DeleteTagResponseObject, error) {
	e := s.client.DeleteTag(int(r.Identifier))

	if not_found.Is(e) {
		return server.DeleteTag404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteTag500JSONResponse(*s.answer(e)), nil
	}

	return server.DeleteTag200Response{}, nil
}
