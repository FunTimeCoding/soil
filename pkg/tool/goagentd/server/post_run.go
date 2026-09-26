package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/generated/server"
)

func (s *Server) PostRun(
	_ context.Context,
	r server.PostRunRequestObject,
) (server.PostRunResponseObject, error) {
	if r.Body.Intent == "" {
		return server.PostRun400JSONResponse{Error: "intent is required"}, nil
	}

	if !s.runner.Start(r.Body.Intent) {
		return server.PostRun409JSONResponse{Error: "agent already running"}, nil
	}

	return server.PostRun202Response{}, nil
}
