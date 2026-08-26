package server

import "github.com/funtimecoding/soil/pkg/tool/goprocessd/process"

func (s *Server) spawn(p *process.Process) {
	s.countMutex.Lock()
	s.running++
	s.countMutex.Unlock()

	if p.Spawn(s.environment.Build(), s.processExited) {
		return
	}

	s.countMutex.Lock()
	s.running--
	s.countMutex.Unlock()
}
