package server

import "github.com/funtimecoding/soil/pkg/strings/join"

func (s *Server) handleList() string {
	names := make([]string, len(s.processes))

	for i, p := range s.processes {
		names[i] = p.Name
	}

	return join.NewLine(names)
}
