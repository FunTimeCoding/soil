package callback

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/web"
)

func (s *Server) Start() {
	if s.verbose {
		console.Format("callback server running on port %d\n", s.port)
	}

	web.ServeAsynchronous(s.server)
}
