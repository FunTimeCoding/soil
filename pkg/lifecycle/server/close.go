package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (s *Server) Close() {
	web.GracefulShutdown(context.Background(), s.http, s.verbose)

	if s.listener != nil {
		errors.LogReclose(s.listener)
	}
}
