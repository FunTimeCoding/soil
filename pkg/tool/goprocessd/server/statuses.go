package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/status"
	"time"
)

func (s *Server) Statuses() []*status.Status {
	processes := s.snapshotProcesses()
	result := make([]*status.Status, len(processes))

	for i, p := range processes {
		started := ""

		if at := p.StartedAt(); !at.IsZero() {
			started = at.Format(time.RFC3339)
		}

		result[i] = status.New(p.Name, p.Command, p.Running(), started)
	}

	return result
}
