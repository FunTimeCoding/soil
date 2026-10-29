package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/server"
)

func (s *Server) StartProcess(
	_ context.Context,
	r server.StartProcessRequestObject,
) (server.StartProcessResponseObject, error) {
	if e := s.service.Start(r.Name); e != nil {
		return server.StartProcess404JSONResponse{Error: e.Error()}, nil
	}

	return server.StartProcess204Response{}, nil
}
