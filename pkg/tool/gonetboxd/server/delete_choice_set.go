package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) DeleteChoiceSet(
	_ context.Context,
	r server.DeleteChoiceSetRequestObject,
) (server.DeleteChoiceSetResponseObject, error) {
	e := s.client.DeleteChoiceSet(r.Name)

	if not_found.Is(e) {
		return server.DeleteChoiceSet404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.DeleteChoiceSet500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteChoiceSet204Response{}, nil
}
