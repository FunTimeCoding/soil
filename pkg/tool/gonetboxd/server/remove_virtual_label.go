package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) RemoveVirtualLabel(
	_ context.Context,
	r server.RemoveVirtualLabelRequestObject,
) (server.RemoveVirtualLabelResponseObject, error) {
	e := s.service.RemoveVirtualLabel(r.Name, r.Key)

	if not_found.Is(e) {
		return server.RemoveVirtualLabel404JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.RemoveVirtualLabel500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.RemoveVirtualLabel204Response{}, nil
}
