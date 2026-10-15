package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) DeleteCustomField(
	_ context.Context,
	r server.DeleteCustomFieldRequestObject,
) (server.DeleteCustomFieldResponseObject, error) {
	e := s.client.DeleteCustomField(r.Name)

	if not_found.Is(e) {
		return server.DeleteCustomField404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.DeleteCustomField500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteCustomField204Response{}, nil
}
