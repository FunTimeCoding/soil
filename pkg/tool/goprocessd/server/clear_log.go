package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/server"
)

func (s *Server) ClearLog(
	_ context.Context,
	r server.ClearLogRequestObject,
) (server.ClearLogResponseObject, error) {
	if e := s.service.ClearLog(r.Name); e != nil {
		return server.ClearLog404JSONResponse{Error: e.Error()}, nil
	}

	return server.ClearLog204Response{}, nil
}
