package server

import "fmt"

func (s *Server) handleStart(arguments []string) string {
	s.commandMutex.Lock()
	defer s.commandMutex.Unlock()

	for _, name := range arguments {
		p := s.findProcess(name)

		if p == nil {
			return fmt.Sprintf("error: unknown process %s", name)
		}

		s.spawn(p)
	}

	return "ok"
}
