package server

import "github.com/funtimecoding/soil/pkg/tool/goprocessd/process"

func (s *Server) setProcesses(processes []*process.Process) {
	s.processMutex.Lock()
	defer s.processMutex.Unlock()
	s.processes = processes
}
