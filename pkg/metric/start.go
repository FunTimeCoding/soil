package metric

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web"
)

func (s *Server) Start() {
	if s.verbose {
		fmt.Printf("start metric server on %d\n", s.port)
	}

	s.wait.Add(1)
	defer s.wait.Done()
	web.ServeAsynchronous(s.server)
}
