package callback

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web"
)

func (s *Server) Start() {
	if s.verbose {
		fmt.Printf("callback server running on port %d\n", s.port)
	}

	web.ServeAsynchronous(s.server)
}
