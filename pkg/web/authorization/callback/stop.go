package callback

import "github.com/funtimecoding/soil/pkg/web"

func (s *Server) Stop() {
	web.GracefulShutdown(s.context, s.server, s.verbose)
}
