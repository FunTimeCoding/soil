package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web"
)

func (s *Server) Stop(verbose bool) {
	web.GracefulShutdown(context.Background(), s.http, verbose)
}
