package server

import "fmt"

func (s *Server) handleStop(arguments []string) string {
	s.commandMutex.Lock()
	defer s.commandMutex.Unlock()

	for _, name := range arguments {
		p := s.findProcess(name)

		if p == nil {
			return fmt.Sprintf("error: unknown process %s", name)
		}

		if e := p.Stop(); e != nil {
			return fmt.Sprintf("error: %s", e)
		}
	}

	return "ok"
}
