package metric

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web"
)

func (s *Server) Stop() {
	if s.verbose {
		fmt.Println("stop metric server")
	}

	web.GracefulShutdown(s.context, s.server, s.verbose)
}
