package supervisor

import "github.com/funtimecoding/soil/pkg/tool/goprocessd/process"

func (s *Supervisor) setProcesses(processes []*process.Process) {
	s.processMutex.Lock()
	defer s.processMutex.Unlock()
	s.processes = processes
}
