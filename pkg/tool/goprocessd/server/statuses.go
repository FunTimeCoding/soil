package server

import "github.com/funtimecoding/soil/pkg/tool/goprocessd/status"

func (s *Server) Statuses() []*status.Status {
	processes := s.snapshotProcesses()
	result := make([]*status.Status, len(processes))

	for i, p := range processes {
		result[i] = status.New(p.Name, p.Command, p.Running())
	}

	return result
}
