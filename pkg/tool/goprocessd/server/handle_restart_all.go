package server

import "fmt"

func (s *Server) handleRestartAll() string {
	s.commandMutex.Lock()
	defer s.commandMutex.Unlock()

	for _, p := range s.snapshotProcesses() {
		if e := p.Stop(); e != nil {
			return fmt.Sprintf("error: %s", e)
		}

		s.spawn(p)
	}

	return "ok"
}
