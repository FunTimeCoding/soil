package server

import "github.com/funtimecoding/soil/pkg/strings/join"

func (s *Server) handleList() string {
	processes := s.snapshotProcesses()
	names := make([]string, len(processes))

	for i, p := range processes {
		names[i] = p.Name
	}

	return join.NewLine(names)
}
