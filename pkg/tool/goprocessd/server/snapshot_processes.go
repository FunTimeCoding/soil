package server

import "github.com/funtimecoding/soil/pkg/tool/goprocessd/process"

func (s *Server) snapshotProcesses() []*process.Process {
	s.processMutex.RLock()
	defer s.processMutex.RUnlock()
	result := make([]*process.Process, len(s.processes))
	copy(result, s.processes)

	return result
}
